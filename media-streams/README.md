# Twilio Media Stream Example

## API Reference

https://www.twilio.com/docs/voice/twiml/stream

* `start.mediaFormat.encoding`: The encoding of the data in the upcoming payload. Value will always be `audio/x-mulaw`
* `start.mediaFormat.sampleRate`: Sample Rate: The Sample Rate in Hertz of the upcoming audio data. Value is always 8000

## Standard

> Best Practices for Audio Recordings
> 
> The telephony standard is 8-bit PCM mono uLaw with a sampling rate of 8Khz. Since this telephony format is fixed, any audio file uploaded to Twilio will be transcoded to that telephony standard. That standard is bandwidth-limited to the 300Hz - 8Khz audio range and is designed for voice and provides acceptable voice-quality results. This standard isn't suitable for quality music reproduction but will provide minimally acceptable results.

https://support.twilio.com/hc/en-us/articles/223180588-Best-Practices-for-Audio-Recordings

## Code

* https://github.com/go-audio/wav
* https://github.com/corticph/slicewriteseek
* https://github.com/twilio/media-streams
* https://stackoverflow.com/questions/58439005/is-there-any-way-to-save-mulaw-audio-stream-from-twilio-in-a-file
* https://github.com/go-audio/wav/issues/2
* https://gist.github.com/dstotijn/9741aecb2ecccf4786939cb534a6f49a

## Articles

* [Live Transcribing Phone Calls using Twilio Media Streams and Google Speech-to-Text](https://www.twilio.com/blog/live-transcribing-phone-calls-using-twilio-media-streams-and-google-speech-text)
* [Best Practices for Audio Recordings](https://support.twilio.com/hc/en-us/articles/223180588-Best-Practices-for-Audio-Recordings)
* [WriteSeeker on a slice](https://www.reddit.com/r/golang/comments/75oyee/writeseeker_on_a_slice/)
