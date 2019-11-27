using GameBox.Pioneer;
using UnityEngine;

namespace Kongfu
{
    public sealed class RotationControl : Control
    {
        public override void OnInit(ITraitContainer container)
        {
            var m = container.NewMatcher();
            m.HasTrait<Renderer>().HasTag(Renderer.Loaded).HasTrait<Movement>().HasTag(Movement.Inited);
            this.filter = container.GetFilter(this, TupleType.Job, m);
        }

        public override void OnUpdate(ITraitContainer container, float deltaTime)
        {
            if (null != this.filter.Target)
            {
                var e = this.filter.Target;

                var m = e.GetTrait<Movement>();
                var r = e.GetTrait<Renderer>();

                var transform = r.View.transform;
                transform.rotation = Quaternion.Euler(0, m.Move.angle + 90, 0);
            }
        }

        private IEntityFilter filter = null;
    }
}