using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace GameBox.Service.Net
{
    public interface IClient
    {
        event Action OnConnected;
        event Action OnClosed;
        event Action<string> OnError;

        IParser Parser { get; set; }

        int ConnectTimeOut { get;set; }

        void Connect(string ip, int port);
        void Disconnect();
        void Send(byte[] buff);
        void Send(byte[] buff, int offset, int count);
    }
}
