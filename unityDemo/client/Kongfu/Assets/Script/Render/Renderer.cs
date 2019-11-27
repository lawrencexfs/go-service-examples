using GameBox.ObjectPool;
using GameBox.Pioneer;
using UnityEngine;

namespace Kongfu
{
    public sealed class Renderer : Trait
    {
        public const string Loaded = "renderer.loaded";

        public override void Dispose()
        {
            if (null != this.View)
            {
                this.Pool.Drop(this.Path, this.View);
                this.View = null;
            }

            this.Pool = null;
            this.Path = null;

            base.Dispose();
        }

        public string Path { get; set; }

        public GameObject View { get; set; }

        public IRecyclePool<GameObject> Pool { get; set; }
    }
}
