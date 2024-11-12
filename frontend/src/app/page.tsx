import Link from 'next/link'

export default function Home() {
  return (
    <>
    <h1>タイトル</h1>
    <div>
    <Link href="/top">
      <button>TOPページへ</button>
    </Link>
    </div>
    <div>
    <Link href = "/users">
      <button>ユーザ管理画面へ</button>
    </Link>
    </div>
    </>
  )
}
