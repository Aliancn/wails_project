export namespace controller {
	
	export class CountDown {
	    duration: number;
	    content: string;
	    deadTime: number;
	    active: boolean;
	    title: string;
	    id: number;
	
	    static createFrom(source: any = {}) {
	        return new CountDown(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.duration = source["duration"];
	        this.content = source["content"];
	        this.deadTime = source["deadTime"];
	        this.active = source["active"];
	        this.title = source["title"];
	        this.id = source["id"];
	    }
	}
	export class TodoList {
	    id: number;
	    taskName: string;
	    competed: boolean;
	
	    static createFrom(source: any = {}) {
	        return new TodoList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.taskName = source["taskName"];
	        this.competed = source["competed"];
	    }
	}
	export class User {
	    username: string;
	    password: string;
	    email: string;
	    avatar: string;
	    id: number;
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.username = source["username"];
	        this.password = source["password"];
	        this.email = source["email"];
	        this.avatar = source["avatar"];
	        this.id = source["id"];
	        this.description = source["description"];
	    }
	}

}

export namespace utils {
	
	export class Resp {
	    code: number;
	    msg: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new Resp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.msg = source["msg"];
	        this.data = source["data"];
	    }
	}

}

