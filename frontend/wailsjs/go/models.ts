export namespace main {
	
	export class DownloadFileResponse {
	    Done: boolean;
	    Skipped: boolean;
	
	    static createFrom(source: any = {}) {
	        return new DownloadFileResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Done = source["Done"];
	        this.Skipped = source["Skipped"];
	    }
	}

}

export namespace types {
	
	export class ImageResponse {
	    Base64string: string;
	    Orientation: string;
	
	    static createFrom(source: any = {}) {
	        return new ImageResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Base64string = source["Base64string"];
	        this.Orientation = source["Orientation"];
	    }
	}

}

