using GameBox;
using GameBox.API.AssetDepot;
using GameBox.Facade;
using GameBox.ObjectPool;
using System;
using UnityEngine;

namespace Kongfu
{
    public class PrefabFactory : IRecycleProcesser<GameObject>, IRecycleFactory<GameObject>
    {
        private RefSet refSet = null;

        public PrefabFactory()
        {
            this.refSet = ResourcesManager.GetUnmanagedRefSet("prefab.factory");
        }

        public void Recover(GameObject obj)
        {
        }

        public void Reclaim(GameObject obj)
        {
            obj.transform.SetParent(null);
            obj.transform.position = new Vector3(0, 10000, 0);
        }

        public GameObject Create(string name)
        {
            var asset = ResourcesManager.LoadAsset<GameObject>(name, this.refSet);
            return GameObject.Instantiate(asset);
        }

        public void CreateAsync(string name, Action<GameObject> handler)
        {
            ResourcesManager.LoadAssetAsync(name, this.refSet, r =>
            {
                if (handler != null)
                {
                    handler(GameObject.Instantiate(r.GameObject()));
                }
            });
        }

        public void Destroy(GameObject obj)
        {
            if (obj != null)
            {
                GameObject.DestroyImmediate(obj);
            }
        }
    }
}
