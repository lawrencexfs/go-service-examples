using GameBox.Pioneer;
using System.Collections.Generic;
using UnityEngine;

namespace Kongfu
{
    public sealed class AnimationControl : Control
    {
        public override void OnInit(ITraitContainer container)
        {
            var m = container.NewMatcher();
            m.HasTrait<Player>()
             .HasTrait<Renderer>()
             .HasTag(Renderer.Loaded)
             .HasTrait<Movement>()
             .HasTag(Movement.Inited);
            this.filter = container.GetFilter(this, TupleType.Job, m);
        }

        public override void OnUpdate(ITraitContainer container, float deltaTime)
        {
            if (null != this.filter.Target)
            {
                var e = this.filter.Target;
                var m = e.GetTrait<Movement>();
                var player = e.GetTrait<Player>();
                if (player.IsDead)
                {
                    PlayAnimation(e, ActionState.Death);
                }
                else if (0 != m.Move.state || ActionState.Unknown == this.nextState)
                {
                    if ((uint)ActionState.Attack == m.Move.state)
                    {
                        if (PlayAnimation(e, ActionState.Attack, ActionState.Idle))
                        {
                            var effect = e.AddTrait<Effect>();
                            if (!e.HasTag(Effect.Inited))
                            {
                                effect.Path = "Effect/AttackEffect0";
                                effect.Duration = 0.5f;
                            }
                            if (!e.HasTrait<Audio>())
                            {
                                var audio = e.AddTrait<Audio>();
                                audio.Path = ExcelManager.Configuration_audio[2].path;
                            }
                        }
                    }
                    else if (m.Move.state >= (uint)ActionState.Hurt)
                    {
                        PlayAnimation(e, ActionState.Hurt);
                    }
                    else if (null == m.LastMove || (m.LastMove.x == m.Move.x && m.LastMove.y == m.Move.y))
                    {
                        PlayAnimation(e, ActionState.Idle);
                    }
                    else
                    {
                        PlayAnimation(e, ActionState.Run);
                    }
                }

                if (ActionState.Unknown != this.nextState)
                {
                    this.duration -= deltaTime;
                    if (this.duration <= 0)
                    {
                        PlayAnimation(e, this.nextState);
                    }
                }
            }
        }

        private bool PlayAnimation(ITraitContainer e, ActionState state, ActionState nextState = ActionState.Unknown)
        {
            if (!CanPlay(state))
            {
                return false;
            }

            this.state = state;

            string animationName = null;
            if (animations.TryGetValue(state, out animationName))
            {
                var animator = GetAnimator(e);
                if (null != animator)
                {
                    animator.CrossFade(animationName, 0.03f, 0);

                    this.nextState = nextState;
                    if (ActionState.Unknown != nextState)
                    {
                        this.duration = animationDurations[state];
                    }

                    return true;
                }
            }

            return false;
        }

        private bool CanPlay(ActionState state)
        {
            HashSet<ActionState> set = null;
            if (conditions.TryGetValue(state, out set))
            {
                return !set.Contains(this.state);
            }

            return true;
        }

        private Animator GetAnimator(ITraitContainer e)
        {
            if (null == this.animator)
            {
                var r = e.GetTrait<Renderer>();
                if (null != r.View)
                {
                    this.animator = r.View.GetComponent<Animator>();
                }
            }

            return this.animator;
        }

        private IEntityFilter filter = null;
        private Animator animator = null;
        private ActionState state = ActionState.Unknown;
        private ActionState nextState = ActionState.Unknown;
        private float duration = 0;

        private enum ActionState
        {
            Unknown = -1,
            Idle = 0,
            Run = 2,
            Attack = 100,
            Dash = 107,
            Hurt = 200,
            Death
        }

        private static readonly Dictionary<ActionState, string> animations = new Dictionary<ActionState, string>
        {
            { ActionState.Idle,          "Idle" },
            { ActionState.Run,           "Run" },
            { ActionState.Attack,        "Attack0" },
            { ActionState.Hurt,          "Take Damage" },
            { ActionState.Death,         "Death" },
        };

        private static readonly Dictionary<ActionState, float> animationDurations = new Dictionary<ActionState, float>
        {
            { ActionState.Idle,          -1f },
            { ActionState.Run,           -1f },
            { ActionState.Attack,         1f },
            { ActionState.Hurt,          -1f },
            { ActionState.Death,          1f },
        };

        private static readonly Dictionary<ActionState, HashSet<ActionState>> conditions = new Dictionary<ActionState, HashSet<ActionState>>
        {
            { ActionState.Idle, new HashSet<ActionState> {
                                        ActionState.Idle,
                                        ActionState.Death } },
            { ActionState.Run, new HashSet<ActionState> {
                                        ActionState.Run,
                                        ActionState.Death } },
            { ActionState.Attack, new HashSet<ActionState> {
                                        ActionState.Attack,
                                        ActionState.Death } },
            { ActionState.Hurt, new HashSet<ActionState> {
                                        ActionState.Hurt,
                                        ActionState.Death } },
            { ActionState.Death, new HashSet<ActionState> {
                                        ActionState.Death } },
        };
    }
}
