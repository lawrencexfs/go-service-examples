using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

namespace Kongfu
{
    public class BattleControl : ViewControl
    {
        protected override void OnLoad()
        {
            var attackButton = Find("Attack").GetComponent<Button>();
            if (null != attackButton)
            {
                attackButton.onClick.AddListener(() =>
                {
                    Game.Broadcast(MessageDef.Attack);
                });
            }

            var joystick = Find("Joystick").GetComponent<Joystick>();
            if (null != joystick)
            {
                joystick.OnJoystickDragging += OnJoystickDragging;
                joystick.OnJoystickEnd += OnJoystickEnd;
            }

            this.hudTemplate = Find("HUD/Template").gameObject;
            this.hudTemplate.SetActive(false);

            this.mainCamera = Camera.main;

            Game.On(MessageDef.AddPlayer, OnAddPlayer);
            Game.On(MessageDef.RemovePlayer, OnRemovePlayer);
        }

        protected override void OnUnload()
        {
            Game.Off(MessageDef.AddPlayer, OnAddPlayer);
            Game.Off(MessageDef.RemovePlayer, OnRemovePlayer);
        }

        protected override void OnUpdate(float deltaTime)
        {
            foreach (var i in this.huds)
            {
                var hud = i.Value;
                var e = Game.GetEntityById(i.Key);
                var renderer = e.GetTrait<Renderer>();
                if (null != renderer && e.HasTag(Renderer.Loaded))
                {
                    var player = e.GetTrait<Player>();
                    var slider = hud.GetComponentInChildren<Slider>();
                    slider.value = player.Data.curhp / 100f;

                    var position = renderer.View.transform.position + HUDOffset;
                    position = this.mainCamera.WorldToScreenPoint(position);
                    hud.transform.position = position;
                    hud.SetActive(true);
                }
                else
                {
                    hud.SetActive(false);
                }
            }
        }

        private void OnJoystickDragging(float x, float y)
        {
            Game.Broadcast(MessageDef.JoyStickDragging, x, y);
        }

        private void OnJoystickEnd()
        {
            Game.Broadcast(MessageDef.JoyStickStop);
        }

        private void OnAddPlayer(MessageDef message, object[] payload)
        {
            var id = (ulong)payload[0];
            if (!this.huds.ContainsKey(id))
            {
                var go = GameObject.Instantiate<GameObject>(this.hudTemplate);
                go.transform.SetParent(this.hudTemplate.transform.parent);
                go.transform.localScale = Vector3.one;

                var e = Game.GetEntityById(id);
                var name = go.GetComponentInChildren<Text>();
                var player = e.GetTrait<Player>();
                name.text = player.Data.name;

                this.huds.Add(e.Id, go);
            }
        }

        private void OnRemovePlayer(MessageDef message, object[] payload)
        {
            var id = (ulong)payload[0];
            GameObject go = null;
            if (this.huds.TryGetValue(id, out go))
            {
                GameObject.Destroy(go);
                this.huds.Remove(id);
            }
        }

        private GameObject hudTemplate = null;
        private Dictionary<ulong, GameObject> huds = new Dictionary<ulong, GameObject>();
        private Camera mainCamera = null;

        private static readonly Vector3 HUDOffset = new Vector3(0, 1, 0);
    }
}
