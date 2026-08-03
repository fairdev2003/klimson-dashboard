dash() {
  local mode="$1"
  local extra_args="${@:2}"

  case "$mode" in
    frontend)
      (cd ~/Desktop/Projects/klimson-dashboard/frontend && npm run dev)
      ;;
    backend)
      local migrate=false
      local backend_args=()
      
      for arg in "${@:2}"; do
        if [ "$arg" = "--migrate" ]; then
          migrate=true
        else
          backend_args+=("$arg")
        fi
      done

      if [ "$migrate" = true ]; then
        echo "[DASH BACKEND] Starting backend with database migration..."
        (cd ~/Desktop/Projects/klimson-dashboard/backend/cmd && AUTO_MIGRATE=true go run . -- "${backend_args[@]}")
      else
        echo "[DASH BACKEND] Starting backend without database migration..."
        (cd ~/Desktop/Projects/klimson-dashboard/backend/cmd && AUTO_MIGRATE=false go run . -- "${backend_args[@]}")
      fi
      ;;
    *)
      echo "Unknown mode: $mode (use: frontend, backend [--migrate])"
      ;;
  esac
}