using GameBox.Pioneer;
using UnityEngine;

namespace Kongfu
{
    public sealed class PositionControl : Control
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
                this.offset.x = m.Move.x * Server2Client - r.View.transform.position.x;
                this.offset.y = m.Move.y * Server2Client - r.View.transform.position.z;
                var deltaX = this.offset.x * deltaTime * MoveSpeed;
                var deltaZ = this.offset.y * deltaTime * MoveSpeed;
                var position = r.View.transform.position;
                r.View.transform.position = new Vector3(position.x + deltaX, 0, position.z + deltaZ);
            }
        }

        private IEntityFilter filter = null;
        private Vector2 offset = Vector2.zero;

        private const float MoveSpeed = 10f;
        private const float Server2Client = 0.01f;
    }
}