import Timeline from '@components/timeline/Timeline';
import TweetForm from '@components/TweetForm';

export default function HomePage() {
  return (
    <>
      <div className="border-b border-twitter-grey p-4">
        <TweetForm />
      </div>
      <Timeline />
    </>
  );
}
