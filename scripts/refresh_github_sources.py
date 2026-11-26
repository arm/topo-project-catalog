import json
import os
import re
import subprocess
from pathlib import Path


def main():
    path = Path(__file__).resolve().parents[1] / "data/github_sources.json"
    pr_body_path = os.environ.get("PR_BODY_PATH")
    sources = json.loads(path.read_text())
    changed = False
    updates = []
    for source in sources:
        sha = subprocess.check_output(
            ["gh", "api", f"repos/{source['repo']}/commits?per_page=1", "--jq", ".[0].sha"],
            text=True,
        ).strip()
        if not re.fullmatch(r"[0-9a-f]{40}", sha):
            raise ValueError(f"Invalid HEAD SHA for {source['repo']}: {sha}")
        print(f"{source['repo']}: {source['sha']} -> {sha}")
        if source["sha"] != sha:
            source["sha"] = sha
            changed = True
            updates.append(source["repo"])
    if pr_body_path:
        Path(pr_body_path).write_text(
            "Detected updated commits for the following repositories:\n\n"
            + "\n".join(f"- {repo}" for repo in updates)
            + "\n"
        )
    if not changed:
        print("No SHA changes detected; leaving the file unchanged.")
        return
    path.write_text(
        "[\n\t" + ",\n\t".join(json.dumps(source) for source in sources) + "\n]\n"
    )


if __name__ == "__main__":
    main()
