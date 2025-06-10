# Initial API Design: Wallet Management App

This document outlines the initial design for the backend API.

## I. Base URL

All API endpoints will be prefixed with `/api/v1`.

## II. Authentication

### 1. User Registration

*   **Endpoint:** `POST /auth/register`
*   **Request Body:**
    ```json
    {
      "name": "John Doe",
      "email": "john.doe@example.com",
      "password": "securepassword123"
    }
    ```
*   **Response (Success 201 Created):**
    ```json
    {
      "userId": "uuid-string",
      "email": "john.doe@example.com",
      "name": "John Doe"
    }
    ```
*   **Response (Error 400 Bad Request):** Invalid input (e.g., email format, password too short).
*   **Response (Error 409 Conflict):** Email already exists.

### 2. User Login

*   **Endpoint:** `POST /auth/login`
*   **Request Body:**
    ```json
    {
      "email": "john.doe@example.com",
      "password": "securepassword123"
    }
    ```
*   **Response (Success 200 OK):**
    ```json
    {
      "token": "jwt-auth-token",
      "user": {
        "userId": "uuid-string",
        "email": "john.doe@example.com",
        "name": "John Doe"
      }
    }
    ```
*   **Response (Error 401 Unauthorized):** Invalid credentials.

### 3. User Logout (Conceptual)

*   **Endpoint:** `POST /auth/logout`
*   **Details:** This might involve invalidating a server-side session or, for JWT, client-side token deletion is primary. If using refresh tokens, this endpoint could invalidate the refresh token.
*   **Response (Success 200 OK):**
    ```json
    {
      "message": "Logged out successfully"
    }
    ```

## III. Wallets (Accounts)

*Authentication required for all wallet endpoints.*

### 1. Create Wallet/Account

*   **Endpoint:** `POST /wallets`
*   **Request Body:**
    ```json
    {
      "name": "My Savings Account",
      "type": "BANK_ACCOUNT" // e.g., BANK_ACCOUNT, CASH, CREDIT_CARD, DIGITAL_YUAN
      "initialBalance": 1000.50,
      "currency": "CNY" // Default, might be implicit
    }
    ```
*   **Response (Success 201 Created):**
    ```json
    {
      "walletId": "uuid-string",
      "userId": "uuid-string",
      "name": "My Savings Account",
      "type": "BANK_ACCOUNT",
      "balance": 1000.50,
      "currency": "CNY",
      "createdAt": "timestamp"
    }
    ```

### 2. Get All Wallets for User

*   **Endpoint:** `GET /wallets`
*   **Response (Success 200 OK):**
    ```json
    [
      {
        "walletId": "uuid-string-1",
        "name": "My Savings Account",
        "type": "BANK_ACCOUNT",
        "balance": 1000.50,
        "currency": "CNY"
      },
      {
        "walletId": "uuid-string-2",
        "name": "Cash Wallet",
        "type": "CASH",
        "balance": 200.00,
        "currency": "CNY"
      }
    ]
    ```

### 3. Get Wallet by ID

*   **Endpoint:** `GET /wallets/{walletId}`
*   **Response (Success 200 OK):** (Similar to create response, includes transaction history - TBD)
*   **Response (Error 404 Not Found):** Wallet not found or does not belong to user.

### 4. Update Wallet

*   **Endpoint:** `PUT /wallets/{walletId}`
*   **Request Body:** (Fields to update, e.g., name)
    ```json
    {
      "name": "My Primary Savings"
    }
    ```
*   **Response (Success 200 OK):** (Updated wallet object)

### 5. Delete Wallet

*   **Endpoint:** `DELETE /wallets/{walletId}`
*   **Response (Success 204 No Content):**
*   **Note:** Consider soft delete vs. hard delete. If transactions are linked, may prevent deletion if not empty.

## IV. Transactions

*Authentication required for all transaction endpoints.*

### 1. Add Transaction (Income/Expense)

*   **Endpoint:** `POST /transactions`
*   **Request Body:**
    ```json
    {
      "walletId": "uuid-string",
      "type": "EXPENSE", // EXPENSE or INCOME
      "amount": 50.00,
      "categoryId": "uuid-string-category", // Link to a category
      "description": "Lunch with colleagues",
      "transactionDate": "YYYY-MM-DDTHH:mm:ssZ", // ISO 8601 date
      "notes": "Optional notes"
    }
    ```
*   **Response (Success 201 Created):** (Full transaction object)
    ```json
    {
      "transactionId": "uuid-string-transaction",
      "walletId": "uuid-string-wallet",
      "type": "EXPENSE",
      "amount": 50.00,
      "categoryId": "uuid-string-category",
      "categoryName": "Food", // Populated from category
      "description": "Lunch with colleagues",
      "transactionDate": "YYYY-MM-DDTHH:mm:ssZ",
      "notes": "Optional notes",
      "createdAt": "timestamp"
      // Wallet balance should be updated server-side
    }
    ```

### 2. Get Transactions for Wallet

*   **Endpoint:** `GET /wallets/{walletId}/transactions`
*   **Query Parameters:** `?page=1&limit=20&startDate=YYYY-MM-DD&endDate=YYYY-MM-DD&type=EXPENSE`
*   **Response (Success 200 OK):** Paginated list of transactions.

### 3. Get All Transactions for User

*   **Endpoint:** `GET /transactions`
*   **Query Parameters:** Similar to above, but across all user's wallets.
*   **Response (Success 200 OK):** Paginated list of transactions.

### 4. Get Transaction by ID

*   **Endpoint:** `GET /transactions/{transactionId}`
*   **Response (Success 200 OK):** (Full transaction object)

### 5. Update Transaction

*   **Endpoint:** `PUT /transactions/{transactionId}`
*   **Request Body:** (Fields to update)
*   **Response (Success 200 OK):** (Updated transaction object)
*   **Note:** Updating amount will require recalculating wallet balance.

### 6. Delete Transaction

*   **Endpoint:** `DELETE /transactions/{transactionId}`
*   **Response (Success 204 No Content):**
*   **Note:** Deleting will require recalculating wallet balance.


## V. Categories (for Transactions)

*Authentication required.*

### 1. Create Category

*   **Endpoint:** `POST /categories`
*   **Request Body:**
    ```json
    {
      "name": "Groceries",
      "type": "EXPENSE" // EXPENSE or INCOME
    }
    ```
*   **Response (Success 201 Created):**
    ```json
    {
      "categoryId": "uuid-string",
      "name": "Groceries",
      "type": "EXPENSE",
      "userId": "uuid-string" // or global categories
    }
    ```

### 2. Get All Categories for User

*   **Endpoint:** `GET /categories`
*   **Query Parameters:** `?type=EXPENSE` (optional)
*   **Response (Success 200 OK):** List of categories.

*(Further endpoints for Update/Delete Category can be added as needed.)*

## VI. Future: Loans, Budgets

Endpoints for Loan Management and Budgeting will be designed later. This initial design focuses on core wallet and transaction functionality.
