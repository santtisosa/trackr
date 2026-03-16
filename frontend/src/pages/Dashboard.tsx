import { useAuth } from '@/contexts/AuthContext'
import { Button } from '@/components/ui/button'

export default function Dashboard() {
  const { user, logout } = useAuth()

  return (
    <div className="min-h-screen bg-background">
      <header className="border-b px-4 py-3 flex items-center justify-between">
        <h1 className="text-lg font-semibold">Trackr</h1>
        <Button variant="outline" size="sm" onClick={logout}>
          Sign out
        </Button>
      </header>

      <main className="p-4">
        <p className="text-muted-foreground text-sm">
          Signed in as {user?.email}
        </p>
      </main>
    </div>
  )
}
