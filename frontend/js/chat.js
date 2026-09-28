let currentConversationId = null;
let currentUserId = null;
let socket = null;

// Initialize chat page
async function initChat() {
	// Check if user is logged in
	if (!checkAuth()) {
		return;
	}

	// Get current user ID
	currentUserId = parseInt(localStorage.getItem("userId"));
	const username = localStorage.getItem("username");

	// Display username
	document.getElementById("username").textContent = username;

	// Load users list
	await loadUsersList();

	// Setup logout button
	document.getElementById("logoutBtn").addEventListener("click", logout);

	// Setup message input
	const messageInput = document.getElementById("messageInput");
	const sendBtn = document.getElementById("sendBtn");

	sendBtn.addEventListener("click", sendMessage);
	messageInput.addEventListener("keypress", (e) => {
		if (e.key === "Enter") {
			sendMessage();
		}
	});
}

// Load all users
async function loadUsersList() {
	try {
		// For now, we'll simulate this. In Step 7, this endpoint doesn't exist yet
		// We'll create it in Step 7b
		console.log("Users list will be loaded here");

		// Placeholder: Show sample users
		const sampleUsers = [
			{ id: 2, username: "bob_jones", email: "bob@example.com", online: true },
			{ id: 3, username: "charlie_brown", email: "charlie@example.com", online: false },
			{ id: 4, username: "diana_prince", email: "diana@example.com", online: true },
		];

		displayUsersList(sampleUsers);
	} catch (error) {
		console.error("Error loading users:", error);
	}
}

// Display users in sidebar
function displayUsersList(users) {
	const usersList = document.getElementById("usersList");
	usersList.innerHTML = "";

	users.forEach((user) => {
		if (user.id === currentUserId) return; // Don't show current user

		const userItem = document.createElement("div");
		userItem.className = "user-item";
		userItem.innerHTML = `
            <div class="user-name">${user.username}</div>
            <div class="user-status ${user.online ? "online" : "offline"}">
                ${user.online ? "🟢 Online" : "⚫ Offline"}
            </div>
        `;

		userItem.addEventListener("click", () => selectUser(user));
		usersList.appendChild(userItem);
	});
}

// Select a user to chat with
async function selectUser(user) {
	// Update active user in sidebar
	document.querySelectorAll(".user-item").forEach((item) => {
		item.classList.remove("active");
	});
	event.currentTarget.classList.add("active");

	// Load conversation
	try {
		// This will be implemented in Step 8 with WebSocket
		console.log("Selected user:", user.username);

		// Update chat header
		document.getElementById("chatHeader").innerHTML = `<h3>${user.username}</h3>`;

		// Clear messages for now
		document.getElementById("messagesList").innerHTML = `
            <div class="no-chat-selected">
                Chat with ${user.username} will appear here
            </div>
        `;
	} catch (error) {
		console.error("Error selecting user:", error);
	}
}

// Send message
function sendMessage() {
	const input = document.getElementById("messageInput");
	const message = input.value.trim();

	if (!message) return;

	// For now, just clear the input
	// WebSocket implementation will be in Step 8
	input.value = "";

	console.log("Message would be sent:", message);
}

// Initialize when page loads
if (document.readyState === "loading") {
	document.addEventListener("DOMContentLoaded", initChat);
} else {
	initChat();
}