using System;

namespace Kongfu
{
    public interface IWindowManager
    {
        void Open(string path, uint layer = 0, float fadeOutTimer = -1);
        void Open<T>(string path, uint layer = 0, Action<T> resolver = null, float fadeOutTimer = -1) where T : ViewControl, new();
        void Close(string path);
    }
}
