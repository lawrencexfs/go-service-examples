using GameBox.API.AssetDepot;
using GameBox.Facade;
using GameBox.Pioneer;
using System;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;

namespace Kongfu
{
    public sealed class WindowSystem : GameBox.Pioneer.System, IWindowManager
    {
        public void Open(string path, uint layer = 0, float fadeOutTimer = -1)
        {
            IEntity e = null;
            if (!this.windows.TryGetValue(path, out e))
            {
                e = this.container.CreateEntity();
                e.AddTrait<Window>().Path = path;

                this.windows.Add(path, e);
            }

            var w = e.GetTrait<Window>();
            w.Layer = layer;
            w.FadeOutTimer = fadeOutTimer;
        }

        public void Open<T>(string path, uint layer = 0, Action<T> resolver = null, float fadeOutTimer = -1) where T : ViewControl, new()
        {
            IEntity e = null;
            if (!this.windows.TryGetValue(path, out e))
            {
                e = this.container.CreateEntity();
                e.AddTrait<Window>().Path = path;

                this.windows.Add(path, e);
            }

            var w = e.GetTrait<Window>();
            w.Layer = layer;
            w.FadeOutTimer = fadeOutTimer;
            var c = e.AddControl<T>();
            if (null != resolver)
            {
                resolver(c);
            }
        }

        public void Close(string path)
        {
            IEntity e = null;
            if (this.windows.TryGetValue(path, out e))
            {
                e.Dispose();
                this.windows.Remove(path);
            }
        }

        public override void OnInit(IEntityContainer container)
        {
            this.container = container;
            this.root = new GameObject("UI");
            this.refSet = ResourcesManager.GetUnmanagedRefSet("UI");
            GameObject.DontDestroyOnLoad(this.root);
            this.windows = new Dictionary<string, IEntity>();

            var m = container.NewMatcher();
            m.HasTrait<Window>().ExceptTag(Window.Inited);
            this.filter = container.GetFilter(this, TupleType.Job, m);
        }

        public override void OnUpdate(IEntityContainer container, float deltaTime)
        {
            foreach (var e in this.filter.Target)
            {
                var w = e.GetTrait<Window>();
                var asset = ResourcesManager.LoadAsset<GameObject>(w.Path, this.refSet);
                w.View = GameObject.Instantiate<GameObject>(asset);
                w.View.transform.SetParent(this.root.transform);
                ApplyWindowLayout(e);
                e.AddTag(Window.Inited);
            }
        }

        private void ApplyWindowLayout(IEntity e)
        {
            if (null == e)
            {
                return;
            }

            var w = e.GetTrait<Window>();

            var slibings = this.windows
                .Where(p => p.Value != e && p.Value.GetTrait<Window>().Layer <= w.Layer && null != p.Value.GetTrait<Window>().View)
                .OrderBy(p => p.Value.GetTrait<Window>().Layer);
            if (slibings.Count() > 0)
            {
                var slibing = slibings.First().Value;
                var srd = slibing.GetTrait<Window>();
                w.View.transform.SetSiblingIndex(srd.View.transform.GetSiblingIndex() + 1);
            }
            else if (this.windows.Count > 0)
            {
                w.View.transform.SetAsFirstSibling();
            }
        }

        private IEntityContainer container = null;
        private IEntitiesFilter filter = null;
        private GameObject root = null;
        private RefSet refSet = null;
        private Dictionary<string, IEntity> windows = null;
    }
}