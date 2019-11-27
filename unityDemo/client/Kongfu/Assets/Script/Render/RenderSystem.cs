using GameBox;
using GameBox.ObjectPool;
using GameBox.Pioneer;
using UnityEngine;

namespace Kongfu
{
    public sealed class RenderSystem : GameBox.Pioneer.System
    {
        public override void OnInit(IEntityContainer container)
        {
            var factory = new PrefabFactory();
            this.pool = GBox.Make<IRecycleManager>().Create(factory, factory);

            var m = container.NewMatcher();
            m.HasTrait<Renderer>().ExceptTag(Renderer.Loaded);
            this.filter = container.GetFilter(this, TupleType.Job, m);
        }

        public override void OnUpdate(IEntityContainer container, float deltaTime)
        {
            foreach (var i in this.filter.Target)
            {
                var r = i.GetTrait<Renderer>();
                r.Pool = this.pool;
                r.View = pool.Pick(r.Path);
                i.AddTag(Renderer.Loaded);
            }
        }

        private IRecyclePool<GameObject> pool = null;
        private IEntitiesFilter filter = null;
    }
}
