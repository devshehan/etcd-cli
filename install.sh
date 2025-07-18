#!/bin/bash

set -e

echo "Installing ETCD Reader CLI - 18th July 2024"

INSTALL_DIR="/usr/local/bin"
CLI_FOLDER="$INSTALL_DIR/etcd-reader-cli"
WRAPPER="$INSTALL_DIR/etcdreader"

echo "Copying CLI files to $CLI_FOLDER"
chmod -R 755 .
sudo rm -rf "$CLI_FOLDER"
sudo mkdir -p "$CLI_FOLDER"
sudo cp -r . "$CLI_FOLDER"

sudo chmod +x "$CLI_FOLDER/run.sh"

echo "Creating etcdreader command"
sudo tee "$WRAPPER" > /dev/null <<EOF
#!/bin/bash
exec "$CLI_FOLDER/run.sh" "\$@"
EOF

sudo chmod +x "$WRAPPER"

echo "Adding environment variables to ~/.bashrc..."

ENV_BLOCK=$(cat <<EOF

# >>> ETCD Reader CLI ENV - Added on $(date) <<<
export SELF_HOSTED_DOMAIN=git.mytaxi.lk
export PERSONAL_GITLAB_TOKEN=******
export GITLAB_PROJECT_ID=3158
# <<< ETCD Reader CLI ENV <<<
EOF
)

if ! grep -q "ETCD Reader CLI ENV" ~/.bashrc; then
    echo "$ENV_BLOCK" >> ~/.bashrc
    echo "ENV block added to .bashrc"
else
    echo "ENV block already exists. Skipping..."
fi

echo ""
echo "Installation complete!"
echo "Run 'source ~/.bashrc' or restart your terminal."
echo "Use your CLI with: etcdreader [--reset]"
