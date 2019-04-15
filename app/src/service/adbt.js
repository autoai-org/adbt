import axios from 'axios'

class AdbtService {
    constructor() {
        this.endpoint = 'http://127.0.0.1:10591'
    }
    async _get(url) {
        return await axios.get(url)
    }
    async _post(url, data) {
        return await axios.post(url, data)
    }
    async _delete(url, data) {
        return await axios.delete(url, {
            data: data
        })
    }
    async getStatus() {
        return await this._get(this.endpoint + '/status')
    }
    async getJobs() {
        return await this._get(this.endpoint + '/jobs')
    }
    async getJobDetail(identifier) {
        return await this._get(this.endpoint+ '/job/' + identifier)
    }
    async getDetailedJobs(jobs) {
        let self = this
        let requestedJobs = jobs
        if(!jobs || typeof(jobs) === 'undefined') {
            requestedJobs = await this.getJobs()
            requestedJobs = requestedJobs.data
        }
        return await Promise.all(requestedJobs.map(async job => {
            return await self.getJobDetail(job.identifier)
        }))
    }
    async deleteJobs(identifier) {
        return await this._delete(this.endpoint + '/jobs', {
            'identifier': identifier
        })
    }
    async addJobs(uri, period, time, database, name) {
        return await this._post(this.endpoint + '/jobs', {
            'uri': uri,
            'periodicity': period,
            'time': time,
            'database': database,
            'name': name
        })
    }
    async testJob(uri, database, name) {
        return await this._post(this.endpoint + '/databases/test', {
            'uri': uri,
            'name': name,
            'database': database
        })
    }
}

let adbtService = new AdbtService()

export {
    adbtService
}