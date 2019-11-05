/*
* This file is part of the GameBox package.
*
* (c) Giant - MouGuangYi<mouguangyi@ztgame.com> , tanxiaoliang<tanxiaoliang@ztgame.com>
*
* For the full copyright and license information, please view the LICENSE
* file that was distributed with this source code.
*
* Document: http://192.168.150.238/GameBox/help/ 
*/

namespace GameBox.Service
{
    /// <summary>
    /// 消息定义
    /// </summary>
    internal class MsgConst
    {
        public static int ClientVertifyReqMsgID = 1;
        public static int ClientVertifySucceedRetMsgID = 2;
        public static int ClientVertifyFailedRetMsgID = 3;
        public static int HeartBeatMsgID = 4;
        public static int SyncMsgID = 10;
        public static int UserDuplicateLoginNotifyID = 24;
        public static int PropSyncClientID = 31;
        public static int MainEntityGenerateID = 32;
        public static int EnterAOIMsgID = 42;
        public static int LeaveAOIMsgID = 43;
        public static int AOIPosChangeMsgID = 44;
        public static int EnterSpaceMsgID = 45;
        public static int LeaveSpaceMsgID = 46;
        public static int UserMoveMsgID = 47;
        public static int SpaceEntityMsgID = 48;
        public static int EntityPosSetMsgID = 49;
        public static int RPCMsgID = 58;
        public static int SyncClockMsgID = 59;

        public static int SpaceUserConnect = 61;
        public static int SpaceUserConnectSucceedRet = 62;
        public static int SyncUserState = 63;
        public static int AOISyncUserState = 64;
        public static int AdjustUserState = 65;

        public static int EntityAOISMsgID = 66;

        public static int EntityBasePropsMsgID = 67;
        public static int EntityEventID = 68;
        public static int AdjustStateRsp = 69;
        public static int HeartBeatResponseMsgID = 70;
        public static int PingMsgID = 71;
        public static int RealRpcMsgRequest = 1000;
        public static int RealRpcMsgResponse = 1001;

        
        public static int EnterReqMsgID = 1000;
        public static int EnterRespMsgID = 1001;
        public static int ActMsgMsgID = 1002;
        public static int TickMsgMsgID = 1003;
        public static int PlayerDisconnectNotifyID = 1004;
        public static int ReadyReqMsgID = 1006;
        public static int GameBeginMsgID = 1007;
    }
}
