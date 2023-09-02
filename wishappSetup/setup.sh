git clone https://github.com/charmbracelet/wish.git
cd wish
sudo nano /etc/systemd/system/domestic_system.service
"""
[Unit]
Description=Domestic System
After=network.target

[Service]
Type=simple
User=dom
Group=dom
WorkingDirectory=/home/dom/
ExecStart=/usr/bin/dom
Restart=on-failure

[Install]
WantedBy=multi-user.target
"""
