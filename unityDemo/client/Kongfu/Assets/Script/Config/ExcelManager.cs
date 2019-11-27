#define _CLIENT_DEFAULT_LOADER_

using System.Collections.Generic;
using System;
using System.IO;

static public class ExcelManager
{
    public delegate Stream ExcelLoader(string name);
#if _CLIENT_DEFAULT_LOADER_
    static private ExcelLoader loader = delegate (string name)
    {
        UnityEngine.TextAsset asset = UnityEngine.Resources.Load<UnityEngine.TextAsset>(name);
        if (asset != null)
            return new System.IO.MemoryStream(asset.bytes);
        return null;
    };
#else
    static private ExcelLoader loader;
#endif
    static public void InitLoader(ExcelLoader l)
    {
        loader = l;
    }
    static private bool _checkLoader()
    {
        if (loader == null)
        {
            throw new System.Exception("Call InitLoader To Set Loader!");
        }
        return true;
    }

	static private List<_Configuration_client_._audio_> _Configuration_audio_list_;
    static public List<_Configuration_client_._audio_> Configuration_audio_list 
    { 
        private set
        {
            _Configuration_audio_list_ = value; 
        }
        get {return _Configuration_audio_list_;} 
    }
	static private Dictionary<uint, _Configuration_client_._audio_> _Configuration_audio_;
    static public Dictionary<uint, _Configuration_client_._audio_> Configuration_audio 
    { 
        private set
        {
            _Configuration_audio_ = value; 
        }
        get {return _Configuration_audio_;} 
    }

    static private bool Load_Configuration()
    {
        Stream s = loader("Config/Configuration");
        if (s != null)
        {
             _Configuration_client_._Excel_ excel = ProtoBuf.Serializer.Deserialize<_Configuration_client_._Excel_>(s);
             if (excel != null)
             {
				Configuration_audio_list = excel.audioData;
				Configuration_audio = new Dictionary<uint, _Configuration_client_._audio_>();
				foreach (_Configuration_client_._audio_ item in excel.audioData)
				{
					if (Configuration_audio.ContainsKey(item.id)) continue;
					Configuration_audio.Add(item.id, item);
				}

                return true;
            }
        }
        return false;
    }
    static public void LoadAll()
    {
        if (_checkLoader())
        {
			Load_Configuration();
        }
    }

    static public System.Collections.IEnumerator LoadAll_Enum()
    {
        yield return LoadAll_Enum(null);
    }
    static public System.Collections.IEnumerator LoadAll_Enum(Action<float> progress)
    {
        if (_checkLoader())
        {
			Load_Configuration();
            if (progress != null)
                progress.Invoke(1f);
            yield return null;

        }
    }

    static public void Unload()
    {
    
		Configuration_audio_list.Clear();
		Configuration_audio_list = null;
		Configuration_audio.Clear();
		Configuration_audio = null;
	}
}