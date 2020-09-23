using GameBox.Service;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Configuration;
using System.Text;

namespace HelloService
{
    public class MsgId
    {
        public static ushort LoginReqMsgID = 1;
        public static ushort LoginRespMsgID = 2;
    }

    public enum ReturnType : uint
    {
        SUCCESS = 0,
        SERVER_BUSY = 1,
        TOKEN_INVALID = 2,
        TOKEN_PARAM_ERR = 3,
        FAIL_RELOGIN = 4
    }

    // LoginReq 登录请求
    // Client ==> LobbyServer
    public class LoginReq
    {
        public string account;
        public string token;
        public ulong uid;
        public string version;
        public byte[] extdata;
    }

    // LoginResp 登录返回
    // LobbyServer ==> Client
    public class LoginResp
    {
        public ReturnType result;
        public string errstr;
        public byte[] extdata;
    }

    public class HelloService : GBoxService
    {
        public event Action OnHelloConnected;
        public event Action<Exception> OnHelloClosed;
        public event Action<Exception> OnHelloError;

        public event Action OnLoginSuccess;
        public event Action<uint> OnLoginFailed;

        private string token;
        private ulong uid;

        public HelloService()
        {
            this.WhenConnected += HelloService_OnConnected;
            this.WhenClosed += HelloService_OnClosed;
            this.WhenError += HelloService_OnError;

            //send proto
            RegisterMsg<LoginReq>(MsgId.LoginReqMsgID);

            //recv proto
            RegisterMsg<LoginResp>(MsgId.LoginRespMsgID, OnLoginResp);
        }

        public void Login(string net, string ip, int port, string token, ulong uid)
        {
            this.token = token;
            this.uid = uid;

            this.Connect(net, ip, port);
        }

        private void sendLoginMsg(string token, ulong uid)
        {
            var loginreq = new LoginReq();
            loginreq.account = "fenggege";
            loginreq.token = token;
            loginreq.uid = uid;
            loginreq.version = "1.0.0.0";
            try
            {
                Send(loginreq);
            }
            catch (Exception ex)
            {
                throw ex;
            }
        }

        private void OnLoginResp(LoginResp resp)
        {
            if (resp.result != ReturnType.SUCCESS)
            {
                if (OnLoginFailed != null)
                {
                    OnLoginFailed((uint)resp.result);
                }
            }
            else
            {
                if (OnLoginSuccess != null)
                {
                    OnLoginSuccess();
                }
            }
        }

        private void HelloService_OnError(Exception obj)
        {
            if (OnHelloError != null)
            {
                OnHelloError(obj);
            }
        }

        private void HelloService_OnClosed(Exception obj)
        {
            if (OnHelloClosed != null)
            {
                OnHelloClosed(obj);
            }
        }

        private void HelloService_OnConnected()
        {
            if (OnHelloConnected != null)
            {
                OnHelloConnected();
            }

            sendLoginMsg(this.token, this.uid);
        }
    }
}
