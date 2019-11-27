using UnityEngine.UI;
using usercmd;

namespace Kongfu
{
    public sealed class DeathControl : ViewControl
    {
        public MsgDeath Msg { private get; set; }

        protected override void OnLoad()
        {
            var titleLabel = Find("TxtStrikeName").GetComponent<Text>();
            if (null != titleLabel)
            {
                titleLabel.text = "你被" + this.Msg.killName + "击败了";
            }

            var scoreLabel = Find("TxtCurScore").GetComponent<Text>();
            if (null != scoreLabel)
            {
                scoreLabel.text = "" + this.Msg.maxScore;
            }

            var backButton = Find("BtnBack").GetComponent<Button>();
            if (null != backButton)
            {
                backButton.onClick.AddListener(() =>
                {
                    Game.Broadcast(MessageDef.ExitMatch);
                    Close();
                });
            }

            var continueButton = Find("BtnContinue").GetComponent<Button>();
            if (null != continueButton)
            {
                continueButton.onClick.AddListener(() =>
                {
                    Game.Broadcast(MessageDef.Reborn);
                    Close();
                });
            }
        }
    }
}
