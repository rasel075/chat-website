// Configuration
const API_BASE_URL = "http://localhost:8080/api";

// Helper function to make API calls
async function apiCall(endpoint, method = "GET", data = null) {
	const options = {
		method: method,
		headers: {
			"Content-Type": "application/json",
		},
	};

	// Add token if available
	const token = localStorage.getItem("token");
	if (token) {
		options.headers.Authorization = `Bearer ${token}`;
	}

	// Add body if provided
	if (data) {
		options.body = JSON.stringify(data);
	}

	try {
		const response = await fetch(`${API_BASE_URL}${endpoint}`, options);
		
		// If response is not ok, throw error
		if (!response.ok) {
			const errorData = await response.json();
			throw new Error(errorData.error || "API Error");
		}

		return await response.json();
	} catch (error) {
		console.error("API Error:", error);
		throw error;
	}
}

// Register user
async function registerUser(username, email, password) {
	return await apiCall("/auth/register", "POST", {
		username,
		email,
		password,
	});
}

// Login user
async function loginUser(email, password) {
	return await apiCall("/auth/login", "POST", {
		email,
		password,
	});
}

// Get all users
async function getAllUsers() {
	return await apiCall("/users", "GET");
}

// Get user profile
async function getUserProfile() {
	return await apiCall("/profile", "GET");
}

// Get or create conversation
async function getOrCreateConversation(otherUserId) {
	return await apiCall("/conversations", "POST", {
		other_user_id: otherUserId,
	});
}

// Get chat messages
async function getMessages(conversationId) {
	return await apiCall(`/messages/${conversationId}`, "GET");
}