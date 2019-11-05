using GameBox;
using GameBox.Channel;
using GameBox.Network;
using GameBox.Service;
using GameBox.Socket;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading;

namespace HelloService
{
    class Program
    {
        static HelloService helloService;
        static Player helloEntity;

        public static GameBox.IServiceProvider[] ServiceProviders
        {
            get
            {
                return new GameBox.IServiceProvider[]
                {
                    new SocketProvider(),
                    new ChannelProvider(),
                    new NetworkProvider(),
                    new GBoxServiceProvider(),
                };
            }
        }

        static void Main(string[] args)
        {
            var gf = new GBoxFramework();
            gf.Bootstrap();

            foreach(var provider in ServiceProviders)
            {
                gf.Register(provider);
            }
            gf.Init();

            helloService = new HelloService();
            helloService.RegisterEntity<Player>(true);
            helloService.OnEntityCreated += onEntityCreated;
            helloService.OnHelloConnected += HelloService_OnConnected;
            helloService.OnHelloClosed += HelloService_OnClosed;
            helloService.OnHelloError += HelloService_OnError;
            helloService.OnLoginSuccess += HelloService_OnLoginSuccess;
            helloService.OnLoginFailed += HelloService_OnLoginFailed;

            helloService.Login("tcp", "127.0.0.1", 17000, "faketoken", 123456);

            while (true)
            {
                helloService.Update(16);
                Thread.Sleep(16);
            }
        }

        static void onEntityCreated(Entity entity)
        {
            helloEntity = entity as Player;

            Console.WriteLine("onEntityCreated id:" + entity.Id);
            //send hello rpc to server
            helloEntity.Rpc("Hello", "wyq", (uint)123);
        }

        static void HelloService_OnLoginFailed(uint obj)
        {
            Console.WriteLine("HelloService_OnLoginFailed errcode:" + obj);
        }

        static void HelloService_OnLoginSuccess()
        {
            Console.WriteLine("HelloService_OnLoginSuccess");
        }

        static void HelloService_OnError(Exception obj)
        {
            Console.WriteLine("HelloService_OnError:" + obj);
        }

        static void HelloService_OnClosed(Exception obj)
        {
            Console.WriteLine("HelloService_OnClosed");
        }

        static void HelloService_OnConnected()
        {
            Console.WriteLine("HelloService_OnConnected");
        }
    }
}
