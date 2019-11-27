using GameBox.Pioneer;
using UnityEngine;

namespace Kongfu
{
    public sealed class CameraSystem : GameBox.Pioneer.System
    {
        public override void OnInit(IEntityContainer container)
        {
            var m = container.NewMatcher();
            m.HasTag(Player.MainPlayer).HasTrait<Renderer>().HasTag(Renderer.Loaded);
            this.filter = container.GetFilter(this, TupleType.Reactive, m);

            this.follower = Camera.main.GetComponent<CameraFollow>();
        }

        public override void OnUpdate(IEntityContainer container, float deltaTime)
        {
            foreach (var e in this.filter.Target)
            {
                var renderer = e.GetTrait<Renderer>();
                this.follower.SetTarget(renderer.View);

                Game.CloseWindow("UI/Loading");
                Game.OpenWindow<BattleControl>("UI/Battle");
            }
        }

        private IEntitiesFilter filter = null;
        private CameraFollow follower = null;
    }
}
