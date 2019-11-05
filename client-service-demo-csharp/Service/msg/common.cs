using System.Collections;
using System.Collections.Generic;

namespace GameBox.Service.msg
{
    public class Ping { }
    public class Pong { }

    // CallMsg 远程调用消息
    public class RpcMsg
    {
        public ulong FromEntityID; // 来自Entity的ID
        public ulong EntityID; // 目的EntityID
        public byte SType;  // 服务类型
        public ulong SID; // 目标服务ID
        public ulong FromSID; // From Service ID
        public ulong Seq; // 序号
        public string methodeName; // 方法名
        public byte[] data; // 参数
        public bool IsSync; // 是否为同步
        public bool IsFromClient; // 是否来自客户端
    }

    // CallRespMsg 远程调用的返回消息
    public class CallRespMsg {
        public ulong Seq; // 序号
        public string ErrString; // 错误
        public byte[] RetData; // 返回的数据
    }

    // 创建实体通知
    public class CreateEntityNotify {
        public string EntityType; //"Player"为玩家
        public ulong EntityID;
    }
}
