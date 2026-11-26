import json
import re
import subprocess
import sys
from datetime import datetime, timezone
from pathlib import Path


def main():
    path = Path(__file__).resolve().parents[1] / "data/github_sources.json"
    sources = json.loads(path.read_text())
    updated_repos = []
    for source in sources:
        sha = subprocess.check_output(
            ["gh", "api", f"repos/{source['repo']}/commits?per_page=1", "--jq", ".[0].sha"],
            text=True,
        ).strip()
        if not re.fullmatch(r"[0-9a-f]{40}", sha):
            raise ValueError(f"Invalid HEAD SHA for {source['repo']}: {sha}")
        print(f"{source['repo']}: {source['sha']} -> {sha}", file=sys.stderr)
        if source["sha"] != sha:
            source["sha"] = sha
            updated_repos.append(source["repo"])
    print(
        f"Checked at: {datetime.now(timezone.utc):%Y-%m-%d %H:%M:%S UTC}\n\n"
        "Detected updated commits for the following repositories:\n\n"
        + "\n".join(f"- {repo}" for repo in updated_repos)
    )
    if len(updated_repos) == 0:
        print("No SHA changes detected; leaving the file unchanged.", file=sys.stderr)
        return
    path.write_text(
        "[\n\t" + ",\n\t".join(json.dumps(source) for source in sources) + "\n]\n"
    )


if __name__ == "__main__":
    main()
