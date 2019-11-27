using GameBox.Pioneer;
using UnityEngine;
using usercmd;

namespace Kongfu
{
    public class BattleState : State
    {
        public override void OnEnter(StateMachine stateMachine)
        {
            Debug.Log("Enter Battle State.");

            var ctx = stateMachine.GetTrait<Context>();
            LoadMap(ctx);

            this.manager = new BattleManager(ctx);

            ctx.Room.OnPacket((int)MsgTypeCmd.Login, OnLogin);
            ctx.Room.OnPacket((int)MsgTypeCmd.AddPlayer, OnAddPlayer);
            ctx.Room.OnPacket((int)MsgTypeCmd.RemovePlayer, OnRemovePlayer);
            ctx.Room.OnPacket((int)MsgTypeCmd.SceneTCP, OnSceneTCP);
            ctx.Room.OnPacket((int)MsgTypeCmd.SceneUDP, OnSceneUDP);
            ctx.Room.OnPacket((int)MsgTypeCmd.Death, OnDeath);
            ctx.Room.Enter(ctx.IP, ctx.Port);

            Game.On(MessageDef.Reborn, OnReborn);
            Game.On(MessageDef.ExitMatch, OnExitMatch);

            this.exited = false;
        }

        public override void OnExecute(StateMachine stateMachine)
        {
            if (this.exited)
            {
                stateMachine.ChangeState(new LoginState());
            }
        }

        public override void OnExit(StateMachine stateMachine)
        {
            this.manager.Dispose();

            var ctx = stateMachine.GetTrait<Context>();
            ctx.Room.Exit();

            UnloadMap();

            Game.CloseWindow("UI/Battle");
        }

        private void LoadMap(Context ctx)
        {
            this.mapEntity = ctx.CreateEntity();
            var renderer = this.mapEntity.AddTrait<Renderer>();
            renderer.Path = "Map";
        }

        private void UnloadMap()
        {
            if (null != this.mapEntity)
            {
                this.mapEntity.Dispose();
                this.mapEntity = null;
            }
        }

        private void OnLogin(IProto proto)
        {
            var result = proto.ToObject<MsgLoginResult>();
            this.manager.Login(result);
        }

        private void OnAddPlayer(IProto proto)
        {
            var msg = proto.ToObject<MsgAddPlayer>();
            this.manager.AddPlayer(msg.player);
        }

        private void OnRemovePlayer(IProto proto)
        {
            var msg = proto.ToObject<MsgRemovePlayer>();
            this.manager.RemovePlayer(msg.id);
        }

        private void OnSceneTCP(IProto proto)
        {
            var msg = proto.ToObject<MsgSceneTCP>();
            this.manager.OnSceneTCP(msg);
        }

        private void OnSceneUDP(IProto proto)
        {
            var msg = proto.ToObject<MsgSceneUDP>();
            this.manager.OnSceneUDP(msg);
        }

        private void OnDeath(IProto proto)
        {
            var msg = proto.ToObject<MsgDeath>();
            this.manager.OnDeath(msg);
        }

        private void OnReborn(MessageDef message, object[] payload)
        {
            this.manager.Reborn();
        }

        private void OnExitMatch(MessageDef message, object[] payload)
        {
            this.manager.Exit();
            this.exited = true;
        }

        private IEntity mapEntity = null;
        private BattleManager manager = null;
        private bool exited = false;
    }
}
