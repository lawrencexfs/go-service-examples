using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace GameBox.Service.Net
{
    public class Client : IClient
    {
        public int ConnectTimeOut { get; set; }

        public IParser Parser { get; set; }

        public event Action OnClosed;
        public event Action OnConnected;
        public event Action<string> OnError;

        private IClient client;
        private string net;

        public Client(string net)
        {
            this.net = net;
        }

        public void Connect(string ip, int port)
        {
            if (client != null)
            {
                client.OnClosed -= OnClosed;
                client.OnConnected -= OnConnected;
                client.OnError -= OnError;
            }

            if (net == "tcp")
            {
                client = new TcpClient();
                client.Parser = Parser;
                client.OnClosed += OnClosed;
                client.OnConnected += OnConnected;
                client.OnError += OnError;
                client.ConnectTimeOut = ConnectTimeOut;
                client.Connect(ip, port);
            }
        }

        public void Disconnect()
        {
            if(client != null)
                client.Disconnect();
        }

        public void Send(byte[] buff)
        {
            if (client != null)
                client.Send(buff);
        }

        public void Send(byte[] buff, int offset, int count)
        {
            if (client != null)
                client.Send(buff, offset, count);
        }
    }
}
