/*
* This file is part of the GameBox package.
*
* (c) Giant - JiangPengQing<jiangpengqing@ztgame.com>
*
* For the full copyright and license information, please view the LICENSE
* file that was distributed with this source code.
*
* Document: http://docs.ztgame.com/asset-depot/
*/

using System.Collections.Generic;
using System.Linq;
using UnityEngine;

namespace GameBox.AssetDepot.Unity
{
    [DisallowMultipleComponent]
    [AddComponentMenu("GameBox/AssetDepot")]
    public class AssetDepotProvider : MonoBehaviour, IServiceProvider
    {
        [SerializeField]
        bool m_simulateABModel = false;

        [SerializeField]
        bool m_syncLoadGameObject = true;

        [SerializeField]
        bool m_asyncLoadGameObject = true;

        IServiceProvider[] m_subProviders;

        public void Register()
        {
            foreach (var p in SubProviders)
                p.Register();
        }

        public void Init()
        {
            foreach (var p in SubProviders)
                p.Init();
        }

        IEnumerable<IServiceProvider> ServiceProviders()
        {
#if UNITY_EDITOR
            if (m_simulateABModel)
                //AB模式，会从ab文件中加载资源，需要在Application.StreamingAssets或者
                //Application.PersistentData目录下存放ab资源文件
                yield return new GameBox.AssetDepot.AssetDepotProvider();
            else
                //正常开发模式时，使用的provider，它会使用AssetDatabase加载资源
                yield return new GameBox.AssetDepot.EditorRuntimeSupport.EditorAssetDepotProvider();
#else
            yield return new GameBox.AssetDepot.AssetDepotProvider();
#endif
            if (m_syncLoadGameObject)
                //同步加载GameObject实例服务
                yield return new SyncGameObjectInstanceLoaderProvider();

            if (m_asyncLoadGameObject)
                //异步加载GameObject实例服务
                yield return new AsyncGameObjectInstanceLoaderProvider();

            //shader管理服务
            yield return new ShaderManagerProvider();
        }

        IServiceProvider[] SubProviders
        {
            get
            {
                if (m_subProviders == null)
                    m_subProviders = ServiceProviders().ToArray();

                return m_subProviders;
            }
        }


    }
}