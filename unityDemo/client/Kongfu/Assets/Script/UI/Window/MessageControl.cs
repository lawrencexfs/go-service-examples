using UnityEngine.UI;

namespace Kongfu
{
    public sealed class MessageControl : ViewControl
    {
        public string Message { private get; set; }

        protected override void OnLoad()
        {
            var text = Find("bg/msg").GetComponent<Text>();
            if (null != text)
            {
                text.text = this.Message;
            }
        }
    }
}
