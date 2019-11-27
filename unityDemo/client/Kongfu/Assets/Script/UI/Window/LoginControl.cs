using System.Collections.Generic;
using UnityEngine.UI;

namespace Kongfu
{
    public class LoginControl : ViewControl
    {
        protected override void OnLoad()
        {
            var serverDropdown = Find("serverList").GetComponent<Dropdown>();
            if (null != serverDropdown)
            {
                serverDropdown.options = ServerOptions;
            }

            var loginButton = Find("login").GetComponent<Button>();
            if (null != loginButton)
            {
                loginButton.onClick.AddListener(() =>
                {
                    var option = serverDropdown.options[serverDropdown.value] as ServerOption;
                    Game.Broadcast(MessageDef.Login, option.IP, option.Port);
                });
            }
        }

        private class ServerOption : Dropdown.OptionData
        {
            public string IP;
            public int Port;
        }

        private readonly List<Dropdown.OptionData> ServerOptions = new List<Dropdown.OptionData>
        {
            new ServerOption
            {
                text = "本机",
                IP = "127.0.0.1",
                Port = 9001,
            },
        };
    }
}
