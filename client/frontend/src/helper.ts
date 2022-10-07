import User from './models/User'
import UserDict from './models/UserDict'
import { useUserStore } from './stores/userStore'

const userStore = useUserStore()

// @devs: functions are sorted by name

export function buttonColor(selectedUser: User, user: User): string {
    if (selectedUser === user) {
        return "is-link"
    } else if (user.Balance < user.MaxDebt) {
        return "is-danger"
    } else if ((user.MaxDebt - user.Balance) > (user.MaxDebt * 0.1)) {
        return "is-warning"
    } else {
        return ""
    }
}

export function displayUserName(user: User): string {
    if (user === undefined) {
        return "???"
    } else
        if (user.BierName !== "") {
            return user.BierName
        } else if (user.LastName !== "" && user.FirstName !== "") {
            return user.FirstName + " " + user.LastName[0] + "."
        } else if (user.FirstName !== "") {
            return user.FirstName
        } else {
            return "???"
        }
}

export function displayUserNameFull(user: User): string {
    let name = ""
    if (user === undefined) {
        return "???"
    } else {
        if (user.FirstName !== "") {
            name = name.concat(user.FirstName)
        }
        if (user.BierName !== "") {
            name = name.concat(" (", user.BierName, ")")
        }
        if (user.LastName !== "") {
            name = name.concat(" ", user.LastName)
        }
    }
    return name
}

export function getUsersAsDict(users: User[]): UserDict {
    const dict = {} as UserDict
    users.forEach(user => {
        let char = ""
        if (user.BierName !== "") {
            char = user.BierName[0].toUpperCase()
        } else {
            char = user.FirstName[0].toUpperCase()
        }
        const charCode = char.charCodeAt(0)
        if (charCode >= 65 && charCode <= 90) { // A-Z
        } else if (charCode >= 48 && charCode <= 57) { // 0-9
            char = "#"
        } else {
            char = "?"
        }
        if (dict[char] === undefined) {
            dict[char] = [user]
        } else {
            dict[char].push(user)
        }
    })
    return dict
}

export function selectUser(user: User): void {
    userStore.selectUser(user)
}

export function sortUserList(users: User[]): User[] {
    users.sort((a, b) => {
        if (displayUserName(a).toLowerCase() < displayUserName(b).toLowerCase()) return -1
        if (displayUserName(a).toLowerCase() > displayUserName(b).toLowerCase()) return 1
        return 0
    })
    return users
}
