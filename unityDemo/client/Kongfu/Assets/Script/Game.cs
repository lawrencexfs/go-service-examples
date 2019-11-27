using GameBox;
using GameBox.Facade;
using GameBox.Pioneer;
using System;
using System.IO;
using UnityEngine;

namespace Kongfu
{
    public sealed class Game
    {
        public static void Start()
        {
            ResourcesManager.InitDepot(null, () =>
            {
                Instance.OnStart();
            });
        }

        public static void Destroy()
        {
            Instance.OnDestroy();
        }

        public static void Update()
        {
            Instance.OnUpdate();
        }

        public static void Broadcast(MessageDef type, params object[] payload)
        {
            Instance.OnBroadcast(type, payload);
        }

        public static bool On(MessageDef type, Action<MessageDef, object[]> handler)
        {
            return Instance.observer.On(type, handler);
        }

        public static bool Off(MessageDef type, Action<MessageDef, object[]> handler)
        {
            return Instance.observer.Off(type, handler);
        }

        public static void OpenWindow(string path, uint layer = 0, float fadeOutTimer = -1)
        {
            Instance.windowManager.Open(path, layer, fadeOutTimer);
        }

        public static void OpenWindow<T>(string path, uint layer = 0, Action<T> resolver = null, float fadeOutTimer = -1) where T : ViewControl, new()
        {
            Instance.windowManager.Open(path, layer, resolver, fadeOutTimer);
        }

        public static void CloseWindow(string path)
        {
            Instance.windowManager.Close(path);
        }

        public static IEntity GetEntityById(ulong id)
        {
            return Instance.world.GetEntityById(id);
        }

        private static Game Instance
        {
            get
            {
                if (null == instance)
                {
                    instance = new Game();
                }

                return instance;
            }
        }

        private static Game instance = null;

        private Game()
        {
        }

        private void OnStart()
        {
            ExcelManager.InitLoader(name =>
            {
                Stream stream = null;
                var refSet = ResourcesManager.GetUnmanagedRefSet("configuration");
                {
                    var asset = ResourcesManager.LoadAsset<TextAsset>(name, refSet);
                    stream = new System.IO.MemoryStream(asset.bytes);
                }
                refSet.Release();

                return stream;
            });
            ExcelManager.LoadAll();

            this.world = GBox.Make<IWorld>();

            this.observer = this.world.AddSystem<MessageSystem>();
            this.windowManager = this.world.AddSystem<WindowSystem>();
            this.world.AddSystem<RenderSystem>();
            this.world.AddSystem<CameraSystem>();
            var room = this.world.AddSystem<RoomSystem>();
            this.world.AddSystem<AudioSystem>();

            this.game = this.world.CreateEntity();
            var ctx = this.game.AddTrait<Context>();
            ctx.World = this.world;
            ctx.Room = room;
            var machine = this.game.AddControl<StateMachine>();
            machine.ChangeState(new LoginState());
        }

        private void OnDestroy()
        {
            if (null != this.world)
            {
                this.world.Dispose();
                this.world = null;
            }

            // 直接赋值为null，因为world的Dispose会清理
            this.observer = null;
            this.windowManager = null;
            this.game = null;
        }

        private void OnUpdate()
        {
            if (null != this.world)
            {
                this.world.Update(Time.deltaTime);
            }
        }

        private void OnBroadcast(MessageDef type, params object[] payload)
        {
            var e = this.world.CreateEntity();
            var msg = e.AddTrait<Message>();
            msg.Type = type;
            msg.Payload = payload;
        }

        private IWorld world = null;
        private IMessageObserver observer = null;
        private IWindowManager windowManager = null;
        private IEntity game = null;
    }
}
