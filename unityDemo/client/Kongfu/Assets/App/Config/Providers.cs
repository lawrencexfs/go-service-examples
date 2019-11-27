/*
 * This file is part of the GameBox package.
 *
 * (c) Giant - MouGuangYi <mouguangyi@ztgame.com> , Yu Bin <yubin@ztgame.com>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Document: http://docs.ztgame.com/gamebox-core
 */

using GameBox;
using GameBox.AssetDepot;
using GameBox.Channel;
using GameBox.MonoDriver;
using GameBox.Network;
using GameBox.ObjectPool;
using GameBox.Pioneer;
using GameBox.Socket;

namespace Game
{
    /// <summary>
    /// 项目注册的服务提供者
    /// </summary>
    public class Providers
    {
        /// <summary>
        /// 项目注册的服务提供者
        /// </summary>
        public static IServiceProvider[] ServiceProviders
        {
            get
            {
                return new IServiceProvider[]
                {
                    // todo: 在此处增加您项目的服务提供者
                    new PioneerProvider(),
                    new MonoDriverProvider(),
                    new AssetDepotProvider(),
                    new ObjectPoolProvider(),
                    new SocketProvider(),
                    new ChannelProvider(),
                    new NetworkProvider(),
                };
            }
        }
    }
}