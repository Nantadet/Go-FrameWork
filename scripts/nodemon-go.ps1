$ErrorActionPreference = "Stop"

$port = 3000

$connections = Get-NetTCPConnection -LocalPort $port -ErrorAction SilentlyContinue

if ($connections) {
    $pids = $connections | Select-Object -ExpandProperty OwningProcess -Unique

    foreach ($processId in $pids) {
        if ($processId -and $processId -ne $PID) {
            Stop-Process -Id $processId -Force -ErrorAction SilentlyContinue
        }
    }

    Start-Sleep -Milliseconds 500
}

go run main.go