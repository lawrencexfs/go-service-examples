using GameBox.Pioneer;
using UnityEngine;

namespace Kongfu
{
    public sealed class EffectControl : Control
    {
        public override void OnInit(ITraitContainer container)
        {
            var initMatcher = container.NewMatcher();
            initMatcher.HasTrait<Effect>()
                       .ExceptTag(Effect.Inited)
                       .HasTrait<Renderer>()
                       .HasTag(Renderer.Loaded);
            this.initFilter = container.GetFilter(this, TupleType.Job, initMatcher);

            var updateMatcher = container.NewMatcher();
            updateMatcher.HasTrait<Effect>()
                         .HasTag(Effect.Inited);
            this.updateFilter = container.GetFilter(this, TupleType.Job, updateMatcher);
        }

        public override void OnUpdate(ITraitContainer container, float deltaTime)
        {
            if (null != this.initFilter.Target)
            {
                var e = this.initFilter.Target;

                var r = e.GetTrait<Renderer>();

                var effect = e.GetTrait<Effect>();
                effect.View = r.Pool.Pick(effect.Path);
                effect.Pool = r.Pool;

                var transform = effect.View.transform;
                transform.SetParent(r.View.transform);
                transform.localPosition = Vector3.up * EffectHeight;
                var particle = transform.GetComponent<ParticleSystem>();
                particle.Play();

                e.AddTag(Effect.Inited);
            }

            if (null != this.updateFilter.Target)
            {
                var e = this.updateFilter.Target;

                var effect = e.GetTrait<Effect>();
                effect.Duration -= deltaTime;
                if (effect.Duration <= 0)
                {
                    e.RemoveTrait<Effect>();
                    e.RemoveTag(Effect.Inited);
                }
            }
        }

        private IEntityFilter initFilter = null;
        private IEntityFilter updateFilter = null;

        private const float EffectHeight = 0.5f;
    }
}