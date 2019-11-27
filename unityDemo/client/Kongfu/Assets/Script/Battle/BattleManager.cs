using GameBox.Pioneer;
using System;
using System.Collections.Generic;
using UnityEngine;
using usercmd;

namespace Kongfu
{
    public sealed class BattleManager : IDisposable
    {
        public BattleManager(Context ctx)
        {
            this.ctx = ctx;
        }

        public void Dispose()
        {
            foreach (var i in this.players)
            {
                i.Value.Dispose();
            }

            this.ctx = null;
            this.mainPlayer = null;
        }

        public void Login(MsgLoginResult msg)
        {
            this.mainPlayerBallId = msg.ballId;

            for (var i = 0; i < msg.others.Count; ++i)
            {
                AddPlayer(msg.others[i]);
            }

            for (var i = 0; i < msg.playerballs.Count; ++i)
            {
                OnPlayerEnterAOI(msg.playerballs[i]);
            }

            for (var i = 0; i < msg.balls.Count; ++i)
            {
                OnItemEnterAOI(msg.balls[i]);
            }
        }

        public void AddPlayer(MsgPlayer msg)
        {
            var e = GetPlayerById(msg.id);
            if (null == e)
            {
                e = this.ctx.CreateEntity();
                e.AddTrait<Player>().Data = msg;

                e.AddTrait<Movement>();

                e.AddControl<RotationControl>();
                e.AddControl<PositionControl>();
                e.AddControl<AnimationControl>();
                e.AddControl<EffectControl>();

                this.players.Add(msg.id, e);

                if (msg.ballId == this.mainPlayerBallId)
                {
                    e.AddControl<InputControl>().Manager = this;
                    e.AddTag(Player.MainPlayer);
                    this.mainPlayer = e;
                }

                Debug.Log(string.Format("{0} enter room.", msg.name));

                Game.Broadcast(MessageDef.AddPlayer, e.Id);
            }
        }

        public void RemovePlayer(ulong id)
        {
            IEntity e = null;
            if (this.players.TryGetValue(id, out e))
            {
                Game.Broadcast(MessageDef.RemovePlayer, e.Id);

                e.Dispose();
                this.players.Remove(id);
            }
        }

        public IEntity GetPlayerById(ulong id)
        {
            IEntity e = null;
            this.players.TryGetValue(id, out e);

            return e;
        }

        public IEntity GetPlayerByBallId(uint ballId)
        {
            foreach (var i in this.players)
            {
                var player = i.Value.GetTrait<Player>();
                if (null != player && player.Data.ballId == ballId)
                {
                    return i.Value;
                }
            }

            return null;
        }

        public uint GetClosetPlayerBallId(float range = float.MaxValue)
        {
            var mpm = this.mainPlayer.GetTrait<Movement>();

            uint targetId = 0;
            float minDistance = float.MaxValue;
            foreach (var i in this.players)
            {
                if (i.Value != this.mainPlayer && i.Value.HasTrait<Renderer>())
                {
                    var pm = i.Value.GetTrait<Movement>();
                    var distance = DistanceSquare(mpm.Move, pm.Move);
                    if (distance < minDistance && distance < range)
                    {
                        minDistance = distance;
                        targetId = i.Value.GetTrait<Player>().Data.ballId;
                    }
                }
            }

            return targetId;
        }

        public void OnSceneTCP(MsgSceneTCP msgSceneTCP)
        {
            for (var i = 0; i < msgSceneTCP.addPlayers.Count; ++i)
            {
                OnPlayerEnterAOI(msgSceneTCP.addPlayers[i]);
            }

            for (var i = 0; i < msgSceneTCP.removePlayers.Count; ++i)
            {
                OnPlayerLeaveAOI(msgSceneTCP.removePlayers[i]);
            }

            for (var i = 0; i < msgSceneTCP.adds.Count; ++i)
            {
                OnItemEnterAOI(msgSceneTCP.adds[i]);
            }

            for (var i = 0; i < msgSceneTCP.removes.Count; ++i)
            {
                OnItemLeaveAOI(msgSceneTCP.removes[i]);
            }

            for (var i = 0; i < msgSceneTCP.eats.Count; ++i)
            {
                var eat = msgSceneTCP.eats[i];
                OnItemLeaveAOI(eat.target);
            }

            for (var i = 0; i < msgSceneTCP.hits.Count; ++i)
            {
                var hit = msgSceneTCP.hits[i];
                var e = GetPlayerByBallId(hit.target);
                if (null != e)
                {
                    var info = e.GetTrait<Player>();
                    info.Data.curhp = hit.curHp;
                }
            }
        }

        public void OnSceneUDP(MsgSceneUDP msgSceneUDP)
        {
            for (var i = 0; i < msgSceneUDP.moves.Count; ++i)
            {
                var move = msgSceneUDP.moves[i];
                var e = GetPlayerByBallId(move.id);
                if (null != e)
                {
                    var m = e.GetTrait<Movement>();
                    m.Move = move;
                    e.AddTag(Movement.Inited);
                }
            }
        }

        public void OnDeath(MsgDeath msgDeath)
        {
            var e = GetPlayerById(msgDeath.Id);
            if (null != e)
            {
                var player = e.GetTrait<Player>();
                if (e.HasTag(Player.MainPlayer))
                {
                    Game.OpenWindow<DeathControl>("UI/Death", 0, c => { c.Msg = msgDeath; });
                }
                else
                {
                    Game.OpenWindow<MessageControl>("UI/Message", 0, c => { c.Message = player.Data.name + "被你击败！"; }, 2);
                }
            }
        }

        public void Reborn()
        {
            SendPacket((int)MsgTypeCmd.ReLife);
        }

        public void Exit()
        {
            SendPacket((int)MsgTypeCmd.ActCloseSocket);
        }

        public void SendPacket(int id, object obj = null)
        {
            if (null != this.ctx)
            {
                this.ctx.Room.SendPacket(id, obj);
            }
        }

        private void OnPlayerEnterAOI(MsgPlayerBall msgPlayerBall)
        {
            var e = GetPlayerByBallId(msgPlayerBall.id);
            if (null != e)
            {
                var m = e.GetTrait<Movement>();
                m.Move = new BallMove
                {
                    x = msgPlayerBall.x,
                    y = msgPlayerBall.y,
                    state = 0,
                };
                e.AddTrait<Renderer>().Path = "Character/Player";
            }
        }

        private void OnPlayerLeaveAOI(uint playerBallId)
        {
            var e = GetPlayerByBallId(playerBallId);
            if (null != e)
            {
                e.RemoveTrait<Renderer>();
                e.RemoveTag(Renderer.Loaded);
                e.RemoveTag(Player.Inited);
            }
        }

        private void OnItemEnterAOI(MsgBall msg)
        {
            if (!this.items.ContainsKey(msg.id))
            {
                string path = null;
                switch (msg.type)
                {
                case 1:
                    path = "Item/Energy";
                    break;
                case 5:
                    path = "Item/Sculpture";
                    break;
                default:
                    return;
                }

                var e = this.ctx.CreateEntity();
                e.AddTrait<Item>().Data = msg;
                e.AddTrait<Renderer>().Path = path;
                e.AddControl<ItemControl>();

                this.items.Add(msg.id, e);
            }
        }

        private void OnItemLeaveAOI(uint ballId)
        {
            IEntity e = null;
            if (this.items.TryGetValue(ballId, out e))
            {
                e.Dispose();
                this.items.Remove(ballId);
            }
        }

        private Context ctx = null;
        private uint mainPlayerBallId = 0;
        private IEntity mainPlayer = null;
        private Dictionary<ulong, IEntity> players = new Dictionary<ulong, IEntity>();
        private Dictionary<uint, IEntity> items = new Dictionary<uint, IEntity>();

        private static float DistanceSquare(BallMove m0, BallMove m1)
        {
            return (m0.x - m1.x) * (m0.x - m1.x) + (m0.y - m1.y) * (m0.y - m1.y);
        }
    }
}
