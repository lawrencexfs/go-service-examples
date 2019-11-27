using GameBox.Pioneer;

namespace Kongfu
{
    public sealed class Audio : Trait
    {
        public const string Playing = "audio.playing";

        public string Path { get; set; }

        public float Duration { get; set; }
    }
}
