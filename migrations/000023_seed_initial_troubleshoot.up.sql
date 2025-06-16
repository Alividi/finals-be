-- Insert into tbl_gangguan
INSERT INTO tbl_gangguan (id, nama_gangguan, deskripsi_gangguan) VALUES
(1, 'Packet Loss Problem', 'Troubleshooting intermittent connection or data packet drops'),
(2, 'Signal Degradation', 'Signal strength is weaker than expected or inconsistent'),
(3, 'Obstructed Signal', 'Dishy cannot get a clear view of the sky due to obstructions'),
(4, 'Connection Lost', 'Service connection completely drops or goes offline'),
(5, 'High Latency', 'Ping is high despite apparent working connection'),
(6, 'Bandwidth Drop', 'Internet speeds are lower than expected'),
(7, 'Frequent Reconnects', 'Connection drops and reconnects repeatedly');

-- Insert into tbl_step
INSERT INTO tbl_step (id, gangguan_id, step, step_number) VALUES
-- Packet Loss
(1, 1, 'Check Cables', 1),
(2, 1, 'Restart the Starlink System', 2),
(3, 1, 'Monitor Connection Again', 3),

-- Signal Degradation
(4, 2, 'Inspect Dishy Mounting', 1),
(5, 2, 'Check for Dirt or Snow', 2),

-- Obstructed Signal
(6, 3, 'Reposition the Dish', 1),
(7, 3, 'Trim Obstructing Trees (If Possible)', 2),

-- Connection Lost
(8, 4, 'Check Service Outage', 1),
(9, 4, 'Inspect Dishy (Hardware)', 2),
(10, 4, 'Contact Support', 3),

-- High Latency
(11, 5, 'Minimize Network Load', 1),
(12, 5, 'Reboot Starlink', 2),

-- Bandwidth Drop
(13, 6, 'Verify Speed Test', 1),
(14, 6, 'Disconnect Idle Devices', 2),

-- Frequent Reconnects
(15, 7, 'Physical Inspection', 1),
(16, 7, 'Replace Cables (If Needed)', 2);

-- Insert into tbl_substep
INSERT INTO tbl_substep (id, step_id, substep, gambar, deskripsi) VALUES
-- Packet Loss
(1, 1, 'Inspect Ethernet cable from Dishy to router.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem1-step1-substep1.jpeg', 'Ensure the Ethernet cable is securely connected and not frayed, bent, or broken. Replace if damaged.'),
(2, 1, 'Check power supply cable for damage.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem1-step1-substep2.jpeg', 'Look for cuts, burns, or wear on the power cable. A faulty cable can cause intermittent connection loss.'),
(3, 2, 'Unplug the Starlink power supply from the wall.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem1-step2-substep1.jpeg', 'Fully remove the plug from the outlet to cut power to the Dishy and router.'),
(4, 2, 'Wait 30 seconds, then plug it back in.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem1-step2-substep2.jpeg', 'Waiting ensures all capacitors discharge and the device properly resets before restarting.'),
(5, 3, 'Open Starlink app and watch Ping statistics.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem1-step3-substep1.jpeg', 'Observe the packet loss statistics for a few minutes to confirm stability.'),

-- Signal Degradation
(6, 4, 'Ensure Dishy is securely mounted and not moving.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem2-step1-substep1.jpeg', 'A firm mount prevents misalignment caused by wind or vibrations, which can impact signal.'),
(7, 5, 'Look for dirt, snow, or debris covering the dish.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem2-step2-substep1.jpeg', 'Snow or dirt can block signal reception and should be cleared.'),
(8, 5, 'Gently clean the dish surface if dirty (use soft cloth).', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem2-step2-substep2.jpeg', 'Use a soft, non-abrasive cloth to clean the dish without scratching it.'),

-- Obstructed Signal
(9, 6, 'Use Starlink app to scan for obstructions.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem3-step1-substep2.jpeg', 'The app highlights blocked areas that affect performance.'),
(10, 6, 'Move Dishy to a location with a full clear view of the sky.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem3-step1-substep2.jpeg', 'Relocate to an open space with minimal overhead obstacles.'),
(11, 7, 'Identify trees, poles, or buildings causing obstruction.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem3-step2-substep1.jpeg', 'Visualize the path between the dish and open sky to spot issues.'),
(12, 7, 'Trim or relocate dish away from obstructions.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem3-step2-substep2.jpeg', 'Safely remove obstacles or reposition Dishy for best visibility.'),

-- Connection Lost
(13, 8, 'Open the Starlink app and check for "Service Outage" alerts.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem4-step1-substep1.jpeg', 'Verify if Starlink is reporting an area-wide issue before troubleshooting locally.'),
(14, 9, 'Look at Dishy — if completely dead (no movement, no heat).', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem4-step1-substep2.jpeg', 'A dish with no activity could indicate hardware failure.'),
(15, 9, 'Ensure power is reaching the Dishy (check power brick lights).', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem4-step2-substep1.jpeg', 'No lights may indicate a dead power supply or disconnected cable.'),
(16, 10, 'If still offline after rebooting, open a Starlink support ticket.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem4-step3-substep1.jpeg', 'Submit a ticket with detailed issue descriptions and screenshots.'),

-- High Latency
(17, 11, 'Stop any large downloads, updates, or heavy streaming temporarily.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem5-step1-substep1.jpeg', 'Heavy network activity can increase latency even on healthy connections.'),
(18, 12, 'Power-cycle the Dishy (same steps as Connection Lost).', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem5-step2-substep1.jpeg', 'Rebooting can clear temporary network congestion issues.'),

-- Bandwidth Drop
(19, 13, 'Run a speed test via Starlink app or fast.com.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem6-step1-substep1.jpeg', 'Compare test results to Starlink advertised speeds.'),
(20, 14, 'Disconnect non-essential devices from Wi-Fi.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem6-step2-substep1.jpeg', 'Reducing the number of connected devices frees up available bandwidth.'),

-- Frequent Reconnects
(21, 15, 'Make sure Dishy’s mount is not shaking in strong winds.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem7-step1-substep1.jpeg', 'Dishy movement causes frequent disconnections.'),
(22, 16, 'Try swapping Ethernet cable or power cables if aged or damaged.', 'https://digisatlink-finals.s3.ap-southeast-1.amazonaws.com/solutions/problem7-step1-substep1.jpeg', 'Cables degrade over time and can cause unstable connectivity.');
