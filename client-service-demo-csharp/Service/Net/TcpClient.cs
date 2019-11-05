using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading;

namespace GameBox.Service.Net
{
    public class TcpClient : IClient
    {
        public event Action OnClosed;
        public event Action OnConnected;
        public event Action<string> OnError;

        public IParser Parser { get; set; }

        public int ConnectTimeOut { get;set; }

        private System.Net.Sockets.Socket socket;
        private IReader reader;

        //private bool m_ConnectTimedOut;
        //private Timer connectTimeoutTimer;
        private string ip;
        private int port;
        private List<byte[]> sendList;
        private bool bDisconnect;
        private bool bError;

        //private delegate void AsyncReceiveMethod();

        public TcpClient()
        {
        }

        public void Connect(string ip, int port)
        {
            try
            {
                this.ip = ip;
                this.port = port;

                //var v = new AsyncReceiveMethod(this.asyncConnectAndSend);
                //v.BeginInvoke(new AsyncCallback(_onRecv), null);
                Thread sendThread = new Thread(asyncConnectAndSend);
                sendThread.IsBackground = true;
                sendThread.Start();
            }
            catch (Exception ex)
            {
                fireOnError(ex.ToString());
            }
        }

        public void Disconnect()
        {
            bDisconnect = true;
        }

        public void Send(byte[] buff)
        {
            lock (sendList)
            {
                sendList.Add(buff);
            }
        }

        public void Send(byte[] buff, int offset, int count)
        {
            var tmp = new byte[count];
            Array.Copy(buff, offset, tmp, 0, count);
            lock (sendList)
            {
                sendList.Add(tmp);
            }
        }

        private void fireOnError(string errdesc)
        {
            if (OnError != null)
                OnError.Invoke(errdesc);
        }

        //private void connectTimeoutTimerDelegate(Object stateInfo)
        //{
        //    // for compression debug statisticsConsole.WriteLine("Connect Timeout");
        //    connectTimeoutTimer.Dispose();
        //    m_ConnectTimedOut = true;
        //    //_socket.Close();
        //}

        //private void EndConnect(IAsyncResult ar)
        //{
        //    if (m_ConnectTimedOut)
        //    {
        //        FireOnError("Attempt to connect timed out");
        //    }
        //    else
        //    {
        //        if (OnConnected != null)
        //            OnConnected.Invoke();

        //        Receive();
        //    }
        //}

        private void asyncConnectAndSend()
        {
            try
            {
                bError = false;
                bDisconnect = false;
                sendList = new List<byte[]>();
                socket = new System.Net.Sockets.Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
                socket.NoDelay = true;

                var address = IPAddress.Parse(ip);
                IPEndPoint endPoint = new IPEndPoint(address, port);

                socket.Connect(endPoint);

                reader = new TcpSocketReader(socket);

                if (OnConnected != null)
                    OnConnected.Invoke();

                Thread recvThread = new Thread(asyncReceive);
                recvThread.IsBackground = true;
                recvThread.Start();

                List<byte[]> list = new List<byte[]>();
                List<byte> tmp = new List<byte>();
                //send loop
                while (true)
                {
                    if(bDisconnect)
                        break;
                    lock (sendList)
                    {
                        if (sendList.Count > 0)
                        {
                            int bufflen = 0;

                            for (int i = 0; i < sendList.Count; i++)
                            {
                                if (bufflen + sendList[i].Length > 8192)
                                {
                                    if (tmp.Count > 0)
                                    {
                                        list.Add(tmp.ToArray());
                                        tmp.Clear();
                                    }
                                    list.Add(sendList[i]);
                                    bufflen = 0;
                                }
                                else
                                {
                                    bufflen += sendList[i].Length;
                                    tmp.AddRange(sendList[i]);
                                }
                                //list.AddRange(sendList[i]);
                            }
                            if (tmp.Count > 0)
                            {
                                list.Add(tmp.ToArray());
                                tmp.Clear();
                            }
                            sendList.Clear();
                        }
                    }

                    if (list.Count > 0)
                    {
                        for (int i = 0; i < list.Count; i++)
                        {
                            socket.Send(list[i], SocketFlags.None);
                        }
                        //var buff = list.ToArray();
                        //socket.Send(buff, SocketFlags.None);
                        list.Clear();
                    }
                    else
                        Thread.Sleep(10);
                }

                if (bError)
                {
                    socket.Close();
                }
                else
                {
                    socket.Close();

                    if (OnClosed != null)
                        OnClosed.Invoke();
                }
            }
            catch (Exception ex)
            {
                if(!bError)
                    fireOnError(ex.ToString());
                socket.Close();
            }
        }

        private void asyncReceive()
        {
            if (Parser != null)
            {
                try { 
                    while (true)
                    {
                        Parser.StartParse(reader);
                    }
                }
                catch (Exception ex)
                {
                    if (!bDisconnect)
                    {
                        bError = true;
                        fireOnError(ex.ToString());
                        Disconnect();
                    }
                }
            }
        }
    }
}
