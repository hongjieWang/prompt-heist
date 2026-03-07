#!/usr/bin/env bash
set -euo pipefail

SERVICE="prompt-heist"
CMD="${1:-}"

case "$CMD" in
  start|stop|restart|status)
    sudo systemctl "$CMD" "$SERVICE"
    ;;
  *)
    echo "Usage: $0 {start|stop|restart|status}"
    exit 1
    ;;
esac
