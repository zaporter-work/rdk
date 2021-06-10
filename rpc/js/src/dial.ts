import { grpc } from "@improbable-eng/grpc-web";
import { CallRequest, CallResponse } from "proto/rpc/webrtc/v1/signaling_pb";
import { SignalingService } from "proto/rpc/webrtc/v1/signaling_pb_service";
import { ClientChannel } from "./ClientChannel";
import { newPeerConnectionForClient } from "./peer";

async function signalCall(signalingAddress: string, host: string, sdp: string): Promise<CallResponse> {
	const callRequest = new CallRequest();
	callRequest.setSdp(sdp);
	let pResolve: (value: CallResponse) => void;
	let pReject: (reason?: any) => void;
	const result = new Promise<CallResponse>((resolve, reject) => {
		pResolve = resolve;
		pReject = reject;
	})
	grpc.unary(SignalingService.Call, {
		request: callRequest,
		metadata: {
			'rpc-host': host,
		},
		host: signalingAddress,
		onEnd: (output: grpc.UnaryOutput<CallResponse>) => {
			const { status, statusMessage, message } = output;
			if (status === grpc.Code.OK && message) {
				pResolve(message);
			} else {
				pReject(statusMessage);
			}
		}
	});
	return await result;
}

export async function dial(signalingAddress: string, host: string): Promise<ClientChannel> {
	const { pc, dc } = await newPeerConnectionForClient();

	const encodedSDP = btoa(JSON.stringify(pc.localDescription));
	const callResponse = await signalCall(signalingAddress, host, encodedSDP);
	const remoteSDP = new RTCSessionDescription(JSON.parse(atob(callResponse.getSdp())));
	pc.setRemoteDescription(remoteSDP);

	const cc = new ClientChannel(pc, dc);
	await cc.ready;
	return cc;
}
