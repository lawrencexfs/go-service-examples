using GameBox.ObjectPool;
using GameBox.Pioneer;
using UnityEngine;

namespace Kongfu
{
    public sealed class Effect : Trait
    {
        public const string Inited = "effect.inited";

        public override void Dispose()
        {
            if (null != this.View && null != this.Pool)
            {
                this.Pool.Drop(this.Path, this.View);
            }

            this.Path = null;
            this.Duration = 0;
            this.View = null;
            this.Pool = null;
        }

        public string Path { get; set; }

        public float Duration { get; set; }

        public GameObject View { get; set; }

        public IRecyclePool<GameObject> Pool { get; set; }
    }
}
