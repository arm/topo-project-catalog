import json
import re
import subprocess
from pathlib import Path


def main():
    path = Path(__file__).resolve().parents[2] / "data/github_sources.json"
    sources = json.loads(path.read_text())
    for source in sources:
        sha = subprocess.check_output(
            ["gh", "api", f"repos/{source['repo']}/commits?per_page=1", "--jq", ".[0].sha"],
            text=True,
        ).strip()
        if not re.fullmatch(r"[0-9a-f]{40}", sha):
            raise ValueError(f"Invalid HEAD SHA for {source['repo']}: {sha}")
        print(f"{source['repo']}: {source['sha']} -> {sha}")
        source["sha"] = sha
    path.write_text(
        "[\n\t" + ",\n\t".join(json.dumps(source) for source in sources) + "\n]\n"
    )


if __name__ == "__main__":
    main()
