export default function ProfilePage({ params }: { params: { userId: string } }) {
  return <p>profile page for {params.userId}</p>;
}
