using GameBox.Channel;
using GameBox.Socket;

namespace GameBox.Service
{
    /// <summary>
    /// Cratos服务提供
    /// </summary>
    public sealed class GBoxServiceProvider : ServiceProvider
    {
        /// <summary>
        /// 注册服务
        /// </summary>
        public override void Init()
        {
            var channelExtend = GBox.Make<IChannelManager>();
            channelExtend.Extend("cratos", (nsp) =>
            {
                var socketFactory = GBox.Make<ISocketManager>();
                var socket = socketFactory.Create(nsp);
                return new Channel.Channel(socket, new GaiaFragment());
            });

            MMStreamPool.Init();

            //GBox.Singleton<Room>().Alias<IRoom>();
        }
    }
}
