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

