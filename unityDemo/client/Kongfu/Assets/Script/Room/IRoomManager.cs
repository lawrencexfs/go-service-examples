using System;

namespace Kongfu
{
    public interface IRoomManager
    {
        void Enter(string ip, int port);
        void Exit();
        void SendPacket(int id, object obj = null);
        void OnPacket(int id, Action<IProto> action);
        void OffPacket(int id, Action<IProto> action);
    }
}
