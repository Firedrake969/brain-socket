export const CYCLED = 'CYCLED';
export const AUTORUN_TOGGLED = 'AUTORUN_TOGGLED';
export const DIRECTORY_CHANGED = 'DIRECTORY_CHANGED';

export function step(amount, generateFrames, saveIO) {
    return function(dispatch, getState) {
        let { socket } = getState();
        socket.socket.emit('cycle', amount, generateFrames, saveIO);
    }
}

export function saveConnectionInfo() {
    return function(dispatch, getState ){
        let { socket } = getState();
        socket.socket.emit('historicalConnectionInfo');
    }
}

export function cycled(frames) {
    return { type: CYCLED, frames: frames }
}

export function saveState(saveName) {
    return function(dispatch, getState) {
        let { socket } = getState();
        socket.socket.emit('save', saveName);
    }
}

export function loadState(loadName) {
    return function(dispatch, getState) {
        let { socket } = getState();
        socket.socket.emit('load', loadName);
    }
}

export function toggleAutorun(autorun) {
    return function(dispatch, getState) {
        let { socket } = getState();
        socket.socket.emit('autorun', autorun);
    }
}

export function autorunToggled(autorun) {
    return { type: AUTORUN_TOGGLED, autorun: autorun };
}

export function changeDirectoryName(name) {
    return function(dispatch, getState) {
        let { socket } = getState();
        socket.socket.emit('directoryName', name);
    }
}

export function directoryChanged(directory) {
    return { type: DIRECTORY_CHANGED, directory: directory };
}

export function createNew(x, y, z, hemispheres, randomize) {
    return function(dispatch, getState) {
        let { socket } = getState();
        socket.socket.emit('create', x, y, z, hemispheres, randomize);
    }
}

export function createSensor(state) {
    return function(dispatch, getState) {
        let { socket } = getState();
        socket.socket.emit('createSensor', state.name, state.radius, state.count, state.plane, state.centerX, state.centerY, state.centerZ, state.outputCount);
    }
}