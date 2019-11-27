using GameBox.Pioneer;
using UnityEngine;

namespace Kongfu
{
    public sealed class ItemControl : Control
    {
        public override void OnInit(ITraitContainer container)
        {
            var m = container.NewMatcher();
            m.HasTrait<Renderer>()
             .HasTag(Renderer.Loaded)
             .HasTrait<Item>()
             .ExceptTag(Item.Inited);
            this.filter = container.GetFilter(this, TupleType.Job, m);
        }

        public override void OnUpdate(ITraitContainer container, float deltaTime)
        {
            if (null != this.filter.Target)
            {
                var e = this.filter.Target;

                var item = e.GetTrait<Item>();
                var r = e.GetTrait<Renderer>();
                r.View.transform.position = new Vector3(item.Data.x * Server2Client, 0, item.Data.y * Server2Client);

                e.AddTag(Item.Inited);
            }
        }

        private IEntityFilter filter = null;

        private const float Server2Client = 0.01f;
    }
}