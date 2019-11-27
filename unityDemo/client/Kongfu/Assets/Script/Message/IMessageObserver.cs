using System;

namespace Kongfu
{
    public enum MessageDef
    {
        Login,
        JoyStickDragging,
        JoyStickStop,
        Attack,
        AddPlayer,
        RemovePlayer,
        Reborn,
        ExitMatch,
    }

    public interface IMessageObserver
    {
        bool On(MessageDef type, Action<MessageDef, object[]> handler);
        bool Off(MessageDef type, Action<MessageDef, object[]> handler);
    }
}