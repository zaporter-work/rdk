interface ReadyPeer {
	pc: RTCPeerConnection;
	dc: RTCDataChannel; 
}

export async function newPeerConnectionForClient(): Promise<ReadyPeer> {
	const peerConnection = new RTCPeerConnection({
		iceServers: [
			{
				// TODO(https://github.com/viamrobotics/core/issues/81): Use Viam STUN/TURN.
				urls: 'stun:stun.erdaniels.com'
			}
		]
	});

	let pResolve: (value: ReadyPeer) => void;
	const result = new Promise<ReadyPeer>(resolve => {
		pResolve = resolve;
	})
	const dataChannel = peerConnection.createDataChannel("data", { negotiated: true, id: 0 });
	dataChannel.binaryType = "arraybuffer";
	
	peerConnection.onicecandidate = async event => {
		if (event.candidate !== null) {
			return;
		}
		pResolve({ pc: peerConnection, dc: dataChannel });
	}

	// set up offer
	const offerDesc = await peerConnection.createOffer();
	try {
		peerConnection.setLocalDescription(offerDesc)
	} catch (e) {
		console.error(e);
	}
	return result;
}
