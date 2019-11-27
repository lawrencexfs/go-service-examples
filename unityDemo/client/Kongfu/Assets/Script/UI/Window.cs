using GameBox.Pioneer;
using UnityEngine;

namespace Kongfu
{
    public sealed class Window : Trait
    {
        public const string Inited = "window.inited";

        public override void Dispose()
        {
            if (null != this.View)
            {
                GameObject.Destroy(this.View);
                this.View = null;
            }

            this.Path = null;
            this.Layer = 0;
            this.FadeOutTimer = -1;
        }

        public string Path { get; set; }

        public uint Layer { get; set; }

        public float FadeOutTimer { get; set; }

        public GameObject View { get; set; }
    }
}