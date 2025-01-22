/** 播放器状态 */
export var PlayerStatus;
(function (PlayerStatus) {
    /** 无状态 */
    PlayerStatus[PlayerStatus["None"] = -2] = "None";
    /** 加载失败 */
    PlayerStatus[PlayerStatus["Failed"] = -1] = "Failed";
    /** 加载中 */
    PlayerStatus[PlayerStatus["Loading"] = 0] = "Loading";
    /** 播放中 */
    PlayerStatus[PlayerStatus["Playing"] = 1] = "Playing";
    /** 暂停中 */
    PlayerStatus[PlayerStatus["Paused"] = 2] = "Paused";
})(PlayerStatus || (PlayerStatus = {}));
