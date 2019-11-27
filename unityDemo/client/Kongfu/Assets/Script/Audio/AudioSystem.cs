using GameBox.Facade;
using GameBox.Pioneer;
using UnityEngine;

namespace Kongfu
{
    public sealed class AudioSystem : GameBox.Pioneer.System
    {
        public override void OnInit(IEntityContainer container)
        {
            var mainCamera = Camera.main.gameObject;
            var refSet = ResourcesManager.GetManagedRefSet(mainCamera);

            var audioSource = GetOrAddComponent<AudioSource>(mainCamera);
            audioSource.clip = ResourcesManager.LoadAsset<AudioClip>(ExcelManager.Configuration_audio[1].path, refSet);
            audioSource.Play();

            var playMatcher = container.NewMatcher();
            playMatcher.HasTrait<Audio>().ExceptTag(Audio.Playing).HasTrait<Renderer>().HasTag(Renderer.Loaded);
            this.playFilter = container.GetFilter(this, TupleType.Job, playMatcher);

            var updateMatcher = container.NewMatcher();
            updateMatcher.HasTrait<Audio>().HasTag(Audio.Playing);
            this.updateFilter = container.GetFilter(this, TupleType.Job, updateMatcher);
        }

        public override void OnUpdate(IEntityContainer container, float deltaTime)
        {
            foreach (var e in this.playFilter.Target)
            {
                var audio = e.GetTrait<Audio>();
                var renderer = e.GetTrait<Renderer>();
                var audioSource = GetOrAddComponent<AudioSource>(renderer.View);
                var refSet = ResourcesManager.GetManagedRefSet(renderer.View);
                var clip = ResourcesManager.LoadAsset<AudioClip>(audio.Path, refSet);
                audio.Duration = clip.length;
                audioSource.clip = clip;
                audioSource.Play();

                e.AddTag(Audio.Playing);
            }

            foreach (var e in this.updateFilter.Target)
            {
                var audio = e.GetTrait<Audio>();
                audio.Duration -= deltaTime;
                if (audio.Duration <= 0)
                {
                    e.RemoveTrait<Audio>();
                    e.RemoveTag(Audio.Playing);
                }
            }
        }

        private T GetOrAddComponent<T>(GameObject go) where T : Component
        {
            var c = go.GetComponent<T>();
            if (null == c)
            {
                c = go.AddComponent<T>();
            }

            return c;
        }

        private IEntitiesFilter playFilter = null;
        private IEntitiesFilter updateFilter = null;
    }
}
